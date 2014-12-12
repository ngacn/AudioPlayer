#include <QtWidgets>

#include "warptencli.h"


HttpRequestInput::HttpRequestInput() {
    initialize();
}

HttpRequestInput::HttpRequestInput(QString v_url_str, QString v_http_method) {
    initialize();
    urlStr = v_url_str;
    httpMethod = v_http_method;
}

void HttpRequestInput::initialize() {
    varLayout = NOT_SET;
    urlStr = "";
    httpMethod = "GET";
}

void HttpRequestInput::add_var(QString key, QString value) {
    vars[key] = value;
}

QString WarptenCli::http_attribute_encode(QString attribute_name, QString input) {
    // result structure follows RFC 5987
    bool need_utf_encoding = false;
    QString result = "";
    QByteArray input_c = input.toLocal8Bit();
    char c;
    for (int i = 0; i < input_c.length(); i++) {
        c = input_c.at(i);
        if (c == '\\' || c == '/' || c == '\0' || c < ' ' || c > '~') {
            // ignore and request utf-8 version
            need_utf_encoding = true;
        }
        else if (c == '"') {
            result += "\\\"";
        }
        else {
            result += c;
        }
    }

    if (result.length() == 0) {
        need_utf_encoding = true;
    }

    if (!need_utf_encoding) {
        // return simple version
        return QString("%1=\"%2\"").arg(attribute_name, result);
    }

    QString result_utf8 = "";
    for (int i = 0; i < input_c.length(); i++) {
        c = input_c.at(i);
        if (
                        (c >= '0' && c <= '9')
                        || (c >= 'A' && c <= 'Z')
                        || (c >= 'a' && c <= 'z')
                        ) {
            result_utf8 += c;
        }
        else {
            result_utf8 += "%" + QString::number(static_cast<unsigned char>(input_c.at(i)), 16).toUpper();
        }
    }

    // return enhanced version with UTF-8 support
    return QString("%1=\"%2\"; %1*=utf-8''%3").arg(attribute_name, result, result_utf8);
}

void WarptenCli::execute(HttpRequestInput *input) {

    // reset variables

    QByteArray request_content = "";
    response = "";
    errorType = QNetworkReply::NoError;
    errorStr = "";


    // decide on the variable layout

    if (input->varLayout == NOT_SET) {
        input->varLayout = input->httpMethod == "GET" || input->httpMethod == "HEAD" ? ADDRESS : URL_ENCODED;
    }


    // prepare request content

    if (input->varLayout == ADDRESS || input->varLayout == URL_ENCODED) {
        // variable layout is ADDRESS or URL_ENCODED

        if (input->vars.count() > 0) {
            bool first = true;
            foreach (QString key, input->vars.keys()) {
                if (!first) {
                    request_content.append("&");
                }
                first = false;

                request_content.append(QUrl::toPercentEncoding(key));
                request_content.append("=");
                request_content.append(QUrl::toPercentEncoding(input->vars.value(key)));
            }

            if (input->varLayout == ADDRESS) {
                input->urlStr += "?" + request_content;
                request_content = "";
            }
        }
    }

    // prepare connection

    QNetworkRequest request = QNetworkRequest(QUrl(input->urlStr));
    request.setRawHeader("User-Agent", "Agent name goes here");

    if (input->varLayout == URL_ENCODED) {
        request.setHeader(QNetworkRequest::ContentTypeHeader, "application/x-www-form-urlencoded");
    }

    if (input->httpMethod == "GET") {
        networkManager->get(request);
    }
    else if (input->httpMethod == "POST") {
        networkManager->post(request, request_content);
    }
    else if (input->httpMethod == "PUT") {
        networkManager->put(request, request_content);
    }
    else if (input->httpMethod == "HEAD") {
        networkManager->head(request);
    }
    else if (input->httpMethod == "DELETE") {
        networkManager->deleteResource(request);
    }
    else {
        QBuffer buff(&request_content);
        networkManager->sendCustomRequest(request, input->httpMethod.toLatin1(), &buff);
    }

}

WarptenCli::WarptenCli(QObject *parent) :
        QObject(parent), networkManager(NULL)
{
    qsrand(QDateTime::currentDateTime().toTime_t());
    networkManager = new QNetworkAccessManager(this);

    connect(networkManager, SIGNAL(finished(QNetworkReply*)),
            this, SLOT(handleNetworkData(QNetworkReply*)));

}

void WarptenCli::handleNetworkData(QNetworkReply *networkReply)
{
    // QUrl url = networkReply->url();
    errorType = networkReply->error();
    if (errorType == QNetworkReply::NoError) {
        response = networkReply->readAll();
    } else {
        errorStr = networkReply->errorString();
    }

    networkReply->deleteLater();
    emit on_execution_finished(this);
}

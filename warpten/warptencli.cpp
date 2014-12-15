#include <QtWidgets>

#include "warptencli.h"

void WarptenCli::execute(QString method, QString path, QUrlQuery &data) {

    // reset variables
    response = "";
    errorType = QNetworkReply::NoError;
    errorStr = "";

    // prepare request content
    QUrl url = baseUrl;
    url.setPath(path);
    QByteArray requestContent = data.query(QUrl::FullyEncoded).toUtf8();
    if (method == "GET") {
        url.setQuery(data);
    }

    // prepare connection
    QNetworkRequest request = QNetworkRequest(url);
    request.setRawHeader("User-Agent", "Agent name goes here");

    if (method != "GET") {
        request.setHeader(QNetworkRequest::ContentTypeHeader, "application/x-www-form-urlencoded");
    }

    if (method == "GET") {
        networkManager->get(request);
    }
    else if (method == "PUT") {
        networkManager->put(request, requestContent);
    }
    else if (method == "POST") {
        networkManager->post(request, requestContent);
    }
    else if (method == "DELETE") {
        networkManager->deleteResource(request);
    }
    qDebug() << method + " " + url.toString(QUrl::FullyEncoded) + " " + requestContent;
}

WarptenCli::WarptenCli(QObject *parent) :
        QObject(parent), networkManager(NULL)
{
    qsrand(QDateTime::currentDateTime().toTime_t());
    networkManager = new QNetworkAccessManager(this);
    baseUrl.setScheme("http");
    baseUrl.setHost("127.0.0.1");
    baseUrl.setPort(7478);

    connect(networkManager, SIGNAL(finished(QNetworkReply*)),
            this, SLOT(handleNetworkData(QNetworkReply*)));
}

void WarptenCli::handleNetworkData(QNetworkReply *networkReply)
{
    errorType = networkReply->error();
    if (errorType == QNetworkReply::NoError) {
        response = networkReply->readAll();
    } else {
        errorStr = networkReply->errorString();
    }

    networkReply->deleteLater();
    emit onExecutionFinished(this);
}

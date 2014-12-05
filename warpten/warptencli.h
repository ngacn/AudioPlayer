#ifndef WARPTENCLI_H
#define WARPTENCLI_H

#include <QObject>
#include <QtNetwork>

class QNetworkReply;

enum HttpRequestVarLayout {NOT_SET, ADDRESS, URL_ENCODED, MULTIPART};

class HttpRequestInput {
public:
    QString urlStr;
    QString httpMethod;
    HttpRequestVarLayout varLayout;
    QMap<QString, QString> vars;

    HttpRequestInput();
    HttpRequestInput(QString v_url_str, QString v_http_method);
    void initialize();
    void add_var(QString key, QString value);
};

class WarptenCli : public QObject
{
    Q_OBJECT
public:
    QByteArray response;
    QNetworkReply::NetworkError errorType;
    QString errorStr;

    explicit WarptenCli(QObject *parent = 0);

    QString http_attribute_encode(QString attribute_name, QString input);
    void execute(HttpRequestInput *input);

signals:
    void on_execution_finished(WarptenCli *worker);

public slots:

private slots:
    void handleNetworkData(QNetworkReply *networkReply);

private:
    QNetworkAccessManager *networkManager;
};

#endif // WARPTENCLI_H

#ifndef WARPTENCLI_H
#define WARPTENCLI_H

#include <QObject>
#include <QtNetwork>

class QNetworkReply;

class WarptenCli : public QObject
{
    Q_OBJECT
public:
    QByteArray response;
    QNetworkReply::NetworkError errorType;
    QString errorStr;

    explicit WarptenCli(QObject *parent = 0);

    void execute(QString method, QString path, QUrlQuery &data);

signals:
    void onExecutionFinished(WarptenCli *cli);

public slots:

private slots:
    void handleNetworkData(QNetworkReply *networkReply);

private:
    QNetworkAccessManager *networkManager;
    QUrl baseUrl;
};

#endif // WARPTENCLI_H

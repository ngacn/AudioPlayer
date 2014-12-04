#ifndef WARPTENCLI_H
#define WARPTENCLI_H

#include <QObject>
#include <QtNetwork>

class QNetworkReply;

class WarptenCli : public QObject
{
    Q_OBJECT
public:
    explicit WarptenCli(QObject *parent = 0);

signals:

public slots:
    void getVersion();

private slots:
    void handleNetworkData(QNetworkReply *networkReply);

private:
    QNetworkAccessManager networkManager;
};

#endif // WARPTENCLI_H

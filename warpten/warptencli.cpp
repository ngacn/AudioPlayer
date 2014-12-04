#include <QtWidgets>

#include "warptencli.h"

WarptenCli::WarptenCli(QObject *parent) :
    QObject(parent)
{
    connect(&networkManager, SIGNAL(finished(QNetworkReply*)),
            this, SLOT(handleNetworkData(QNetworkReply*)));

}

void WarptenCli::handleNetworkData(QNetworkReply *networkReply)
{
    QUrl url = networkReply->url();
    if (!networkReply->error()) {
        QByteArray response(networkReply->readAll());
        QMessageBox::information(NULL, tr("Warpten Client"),
                             tr("The following error occurred: %1.")
                             .arg(response.data()));
    }

    networkReply->deleteLater();
}

void WarptenCli::getVersion()
{
    QString url = QString("http://127.0.0.1:7478/version");
    networkManager.get(QNetworkRequest(url));
}

#include <QtWidgets>

#include "playlisttab.h"
#include "warptencli.h"

PlaylistTab::PlaylistTab(const QString &uuid, QWidget *parent) :
        QWidget(parent), uuid(uuid)
{
    tracksListBox = new QListWidget;
    QVBoxLayout *layout = new QVBoxLayout;
    layout->addWidget(tracksListBox);
    setLayout(layout);
}

void PlaylistTab::addTrack(const QString &uuid, const QString &path)
{
    QListWidgetItem *item = new QListWidgetItem;
    item->setData(Qt::DisplayRole, path);
    item->setData(Qt::UserRole, uuid);
    tracksListBox->addItem(item);
}

void PlaylistTab::delTrack(const QString &uuid, int row)
{
    QUrlQuery query;
    query.addQueryItem("uuid", uuid);
    query.addQueryItem("index", QString::number(row));
    WarptenCli *cli = new WarptenCli(this);
    connect(cli, SIGNAL(on_execution_finished(WarptenCli*)), this, SLOT(updateDelTrack(WarptenCli*)));
    cli->execute("POST", "/track/del", query);
}

void PlaylistTab::updateDelTrack(WarptenCli *cli)
{
    if (cli->errorType != QNetworkReply::NoError) {
        // an error occurred
        return;
    }
    QJsonDocument loadDoc(QJsonDocument::fromJson(cli->response));
    QJsonObject json = loadDoc.object();
    if (!json["err"].toString().isEmpty()) {
        qDebug() << json["err"].toString();
        return;
    }
    int row = json["return"].toInt();
    delete tracksListBox->item(row);
}

const QString& PlaylistTab::getUuid()
{
    return uuid;
}

void PlaylistTab::contextMenuEvent(QContextMenuEvent *e)
{
    QPoint p = tracksListBox->mapFromGlobal(e->globalPos());
    QListWidgetItem *item = tracksListBox->itemAt(p);
    if (item) {
        QMenu *contextMenu = new QMenu(tracksListBox);
        QAction *delTrackAct = new QAction(tr("&Delete track"), tracksListBox);
        contextMenu->addAction(delTrackAct);
        QAction *triggeredAct = contextMenu->exec(e->globalPos());
        if (triggeredAct == delTrackAct) {
            delTrack(item->data(Qt::UserRole).toString(), tracksListBox->row(item));
        }
    }
}

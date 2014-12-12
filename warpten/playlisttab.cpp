#include <QtWidgets>

#include "playlisttab.h"

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

const QString& PlaylistTab::getUuid()
{
    return uuid;
}

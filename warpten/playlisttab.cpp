#include <QtWidgets>

#include "playlisttab.h"

PlaylistTab::PlaylistTab(const QString &uuid, QWidget *parent) :
        QWidget(parent), uuid(uuid)
{
    QListWidget *tracksListBox = new QListWidget;
    QStringList tracks;

    for (int i = 1; i <= 30; ++i)
        tracks.append(tr("Track %1").arg(i));
    tracksListBox->insertItems(0, tracks);

    QVBoxLayout *layout = new QVBoxLayout;
    layout->addWidget(tracksListBox);
    setLayout(layout);
}

const QString& PlaylistTab::getUuid()
{
    return uuid;
}

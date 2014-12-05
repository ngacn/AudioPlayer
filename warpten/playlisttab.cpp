#include <QtWidgets>

#include "playlisttab.h"

PlaylistTab::PlaylistTab(QWidget *parent) :
    QWidget(parent)
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

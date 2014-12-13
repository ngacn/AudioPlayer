#ifndef PLAYLISTTAB_H
#define PLAYLISTTAB_H

#include <QWidget>

class QListWidget;

class WarptenCli;

class PlaylistTab : public QWidget
{
    Q_OBJECT
public:
    explicit PlaylistTab(const QString &uuid, QWidget *parent = 0);
    void addTrack(const QString &uuid, const QString &path);
    const QString& getUuid();

signals:

private slots:
    void delTrack(const QString &uuid, int row);
    void updateDelTrack(WarptenCli *cli);

private:
    void contextMenuEvent(QContextMenuEvent *e);

    QString uuid;
    QListWidget *tracksListBox;
};

#endif // PLAYLISTTAB_H

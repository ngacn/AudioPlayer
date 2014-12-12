#ifndef PLAYLISTTAB_H
#define PLAYLISTTAB_H

#include <QWidget>

class PlaylistTab : public QWidget
{
    Q_OBJECT
public:
    explicit PlaylistTab(const QString &uuid, QWidget *parent = 0);
    const QString& getUuid();

signals:

public slots:

private:
    QString uuid;
};

#endif // PLAYLISTTAB_H

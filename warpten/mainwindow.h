#ifndef MAINWINDOW_H
#define MAINWINDOW_H

#include <QMainWindow>

QT_BEGIN_NAMESPACE
class QAction;
class QMenu;
class QProcess;
QT_END_NAMESPACE

class WarptenCli;

class MainWindow : public QMainWindow
{
    Q_OBJECT

public:
    MainWindow();

protected:
    void closeEvent(QCloseEvent *event);

private slots:
    void newTrack();
    void newPlaylist();
    void about();
    void aboutDawn();

    void updateVersion(WarptenCli *cli);
    void updatePlaylists(WarptenCli *cli);
    void updateNewTrack(WarptenCli *cli);
    void updateNewPlaylist(WarptenCli *cli);
    void updateCloseTab(WarptenCli *cli);
    void requestCloseTab(int index);

private:
    void createActions();
    void createMenus();
    void createStatusBar();
    void readSettings();
    void writeSettings();

    void requestVersion();
    void requestPlaylists();

    QMenu *fileMenu;
    QMenu *editMenu;
    QMenu *helpMenu;

    QAction *newTrackAct;
    QAction *newPlaylistAct;
    QAction *exitAct;
    QAction *aboutAct;
    QAction *aboutQtAct;
    QAction *aboutDawnAct;

    QTabWidget *playlistsTabWidget;

    QString version;

    QProcess *daemonProcess;
};

#endif

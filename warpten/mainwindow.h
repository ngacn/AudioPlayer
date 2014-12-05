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
    void newPlaylist();
    void about();

    void updateVersion(WarptenCli *cli);

private:
    void createActions();
    void createMenus();
    void createStatusBar();
    void readSettings();
    void writeSettings();

    void requestVersion();

    QMenu *fileMenu;
    QMenu *editMenu;
    QMenu *helpMenu;
    QAction *newPlaylistAct;
    QAction *exitAct;
    QAction *aboutAct;
    QAction *aboutQtAct;
    QTabWidget *playlistsTabWidget;

    QString version;

    QProcess *daemonProcess;
};

#endif

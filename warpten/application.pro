QT += widgets

HEADERS       = mainwindow.h \
    playlisttab.h
SOURCES       = main.cpp \
                mainwindow.cpp \
    playlisttab.cpp
RESOURCES     = application.qrc

# install
target.path = $$[QT_INSTALL_EXAMPLES]/widgets/mainwindows/application
INSTALLS += target

Screen Router Module
====================

Source file: ``internal/widgets/screenrouter.go``

Overview
--------

This module connects the dashboard widgets to their full-screen pages. The dashboard has five selectable panels, and each panel index maps to a specific page such as Calendar, Courses, Tasks, Schedule, or Assignments.

The router keeps page switching centralized so the main program does not need to know the details of every screen.

Functions
---------

ScreenRouter
~~~~~~~~~~~~

.. code-block:: go

   func ScreenRouter(
       app *tview.Application,
       widgets []tview.Primitive,
       layout *tview.Flex,
       data *model.AppData,
   ) func(index int)

Creates and returns a callback that opens a screen based on a selected dashboard index.

``app``
    The running ``tview.Application``. It is used to replace the root screen.

``widgets``
    The dashboard widgets. These are used to restore focus when returning from a full page.

``layout``
    The main dashboard layout. Sub-pages use this to return to the dashboard.

``data``
    Pointer to the loaded application data. This is passed into each page so the screens can display courses, assignments, tasks, and student information.

How it works:

``ScreenRouter`` returns an anonymous function. When that function receives an index, it creates the matching page and sets it as the root screen.

Index mapping:

``0``
    Opens the Calendar page.

``1``
    Opens the Courses page.

``2``
    Opens the Tasks page.

``3``
    Opens the Schedule page.

``4``
    Opens the Assignments page.

Each page receives a ``returnTo`` callback. That callback restores the dashboard layout and returns focus to the widget that opened the page.

Default case:

If an unknown index is passed in, the router creates a simple fallback ``TextView`` showing which box was selected. Pressing Escape returns to the dashboard.

Notes
-----

The router does not decide how the user navigates. It only opens pages. Keyboard handling is handled separately by the navigation module.

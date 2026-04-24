Navigation Module
=================

Source file: ``internal/widgets/navigation.go``

Overview
--------

This module handles keyboard navigation for the main dashboard. It is focused only on the dashboard layout, not the entire application, so text input pages such as the Canvas settings screen can receive normal keyboard input without being interrupted by dashboard shortcuts.

The dashboard uses five selectable widgets. Pressing number keys changes which widget is selected, pressing Enter opens the selected screen, pressing ``c`` opens Canvas settings, and pressing ``q`` quits the application.

Functions
---------

HandleNavigation
~~~~~~~~~~~~~~~~

.. code-block:: go

   func HandleNavigation(
       dashboard *tview.Flex,
       widgets []tview.Primitive,
       openScreen func(index int),
       openCanvasSettings func(),
       stopApp func(),
   )

Installs keyboard controls on the main dashboard.

``dashboard``
    The root dashboard layout. This is where the input handler is attached.

``widgets``
    The five dashboard panels that can be selected by the user.

``openScreen``
    Callback used when Enter is pressed. It receives the selected widget index and opens the matching full page.

``openCanvasSettings``
    Callback used when the user presses ``c`` or ``C``. It opens the Canvas settings page.

``stopApp``
    Callback used when the user presses ``q`` or ``Q``. In practice, this is usually ``app.Stop``.

How it works:

The function stores the currently selected dashboard panel in ``focusIndex``. It starts at ``0``, which means the first dashboard widget is selected by default.

Inside the function, ``updateFocus`` loops over every dashboard widget and updates the border color, title color, and border style. The selected widget is shown with a bold spring-green border. All other widgets return to a normal white border.

The input handler is attached with ``dashboard.SetInputCapture``. Because the handler belongs only to the dashboard, it does not hijack keys typed into other screens.

Keys to Note:

``1`` through ``5``
    Select one of the five dashboard panels.

``Enter``
    Opens the currently selected panel using ``openScreen``.

``c`` or ``C``
    Opens the Canvas settings page.

``q`` or ``Q``
    Quits the application through ``stopApp``.

Notes
-----

This file is intentionally dashboard-specific. Sub-pages should define their own local input handling, usually using Escape to return to the dashboard.

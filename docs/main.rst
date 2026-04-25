main.go
=======

Purpose
-------

``main.go`` is the entry point for VIO. It loads cached local JSON data, loads optional Canvas configuration, builds the dashboard UI, wires navigation, and starts the terminal application. It also owns the Canvas background sync loop so the app can update courses and assignments without blocking the interface.

Imports and Dependencies
------------------------

This file connects several internal packages:

- ``internal/store`` loads and saves local JSON data and Canvas settings
- ``internal/canvas`` syncs active Canvas courses and assignments
- ``internal/views/settings`` renders the Canvas settings screen
- ``internal/widgets`` builds the dashboard and routes to each page
- ``tview`` provides the terminal UI application and modal widgets

Function Documentation
----------------------

main
~~~~

.. code-block:: go

   func main()

Starts the full VIO application.

How it works:

1. Loads local app data from ``data.json`` with ``store.LoadAppData``
2. Loads Canvas settings with ``store.LoadCanvasConfig``
3. Creates a new ``tview.Application``
4. Builds the dashboard widgets using the cached JSON data.
5. Sets up helper closures for dashboard refresh, Canvas syncing, polling, screen switching, and settings
6. Installs dashboard navigation with ``widgets.HandleNavigation``
7. Shows a first-run Canvas setup modal if no Canvas config exists
8. Starts the TUI event loop with ``app.Run``

Important internal helpers inside ``main``:

``refreshDashboard``
    Calls ``widgets.RefreshMainWidgets`` so dashboard boxes redraw after data changes.

``syncNow``
    Starts a background goroutine that calls ``canvas.SyncCoursesAndAssignments``. On success, it replaces the in-memory data, saves the updated cache with ``store.SaveAppData``, refreshes the dashboard, and updates the Canvas sync status message.

``startPolling``
    Restarts the Canvas polling loop. If valid Canvas credentials exist, it immediately syncs once, then syncs again every configured interval. The default interval is 5 minutes.

``showDashboard``
    Returns the UI to the main dashboard root.

``openCanvasSettings``
    Opens the Canvas settings page. Saving from this page stores the new config, restarts polling, and returns to the dashboard. Canceling returns to the dashboard without changing settings.

Notes
-----

The app loads cached JSON first, then syncs Canvas in the background. This makes startup fast and allows the app to still be usable if Canvas is unavailable.

"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
self["webpackHotUpdate_N_E"]("app/page",{

/***/ "(app-pages-browser)/./app/components/charts/BarChart.tsx":
/*!********************************************!*\
  !*** ./app/components/charts/BarChart.tsx ***!
  \********************************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"(app-pages-browser)/./node_modules/next/dist/compiled/react/jsx-dev-runtime.js\");\n/* harmony import */ var _app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @/app/utils/Utils */ \"(app-pages-browser)/./app/utils/Utils.ts\");\n/* harmony import */ var chart_js__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! chart.js */ \"(app-pages-browser)/./node_modules/chart.js/dist/chart.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! react */ \"(app-pages-browser)/./node_modules/next/dist/compiled/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_2___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_2__);\n/* harmony import */ var _ChartjsConfig__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./ChartjsConfig */ \"(app-pages-browser)/./app/components/charts/ChartjsConfig.tsx\");\n/* __next_internal_client_entry_do_not_use__ default auto */ \nvar _s = $RefreshSig$();\n\n\n\n\nchart_js__WEBPACK_IMPORTED_MODULE_4__.Chart.register(chart_js__WEBPACK_IMPORTED_MODULE_4__.BarController, chart_js__WEBPACK_IMPORTED_MODULE_4__.BarElement, chart_js__WEBPACK_IMPORTED_MODULE_4__.LinearScale, chart_js__WEBPACK_IMPORTED_MODULE_4__.TimeScale, chart_js__WEBPACK_IMPORTED_MODULE_4__.Tooltip, chart_js__WEBPACK_IMPORTED_MODULE_4__.Legend);\nconst BarChart = (param)=>{\n    let { data, width, height } = param;\n    _s();\n    const [chart, setChart] = (0,react__WEBPACK_IMPORTED_MODULE_2__.useState)(null);\n    const canvas = (0,react__WEBPACK_IMPORTED_MODULE_2__.useRef)(null);\n    const legend = (0,react__WEBPACK_IMPORTED_MODULE_2__.useRef)(null);\n    const { textColor, gridColor, tooltipBodyColor, tooltipBgColor, tooltipBorderColor } = _ChartjsConfig__WEBPACK_IMPORTED_MODULE_3__.chartColors;\n    (0,react__WEBPACK_IMPORTED_MODULE_2__.useEffect)(()=>{\n        const ctx = canvas.current;\n        // eslint-disable-next-line no-unused-vars\n        const newChart = new chart_js__WEBPACK_IMPORTED_MODULE_4__.Chart(ctx, {\n            type: \"bar\",\n            data: data,\n            options: {\n                layout: {\n                    padding: {\n                        top: 12,\n                        bottom: 16,\n                        left: 20,\n                        right: 20\n                    }\n                },\n                scales: {\n                    y: {\n                        border: {\n                            display: false\n                        },\n                        ticks: {\n                            maxTicksLimit: 5,\n                            callback: (value)=>(0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.formatValue)(value),\n                            color: textColor.dark\n                        },\n                        grid: {\n                            color: gridColor.dark\n                        }\n                    },\n                    x: {\n                        type: \"time\",\n                        time: {\n                            parser: \"MM-DD-YYYY\",\n                            unit: \"month\",\n                            displayFormats: {\n                                month: \"MMM YY\"\n                            }\n                        },\n                        border: {\n                            display: false\n                        },\n                        grid: {\n                            display: false\n                        },\n                        ticks: {\n                            color: textColor.dark\n                        }\n                    }\n                },\n                plugins: {\n                    legend: {\n                        display: false\n                    },\n                    tooltip: {\n                        callbacks: {\n                            //title: () => false, // Disable tooltip title\n                            label: (context)=>(0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.formatValue)(context.parsed.y)\n                        },\n                        bodyColor: tooltipBodyColor.dark,\n                        backgroundColor: tooltipBgColor.dark,\n                        borderColor: tooltipBorderColor.dark\n                    }\n                },\n                interaction: {\n                    intersect: false,\n                    mode: \"nearest\"\n                },\n                animation: {\n                    duration: 500\n                },\n                maintainAspectRatio: false,\n                resizeDelay: 200\n            },\n            plugins: [\n                {\n                    id: \"htmlLegend\",\n                    afterUpdate (c, args, options) {\n                        const ul = legend.current;\n                        if (!ul) return;\n                        // Remove old legend items\n                        while(ul.firstChild){\n                            ul.firstChild.remove();\n                        }\n                        // Reuse the built-in legendItems generator\n                        const items = c.options.plugins.legend.labels.generateLabels(c);\n                        items.forEach((item)=>{\n                            const li = document.createElement(\"li\");\n                            li.style.marginRight = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.margin[4];\n                            // Button element\n                            const button = document.createElement(\"button\");\n                            button.style.display = \"inline-flex\";\n                            button.style.alignItems = \"center\";\n                            button.style.opacity = item.hidden ? \".3\" : \"\";\n                            button.onclick = ()=>{\n                                c.setDatasetVisibility(item.datasetIndex, !c.isDatasetVisible(item.datasetIndex));\n                                c.update();\n                            };\n                            // Color box\n                            const box = document.createElement(\"span\");\n                            box.style.display = \"block\";\n                            box.style.width = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.width[3];\n                            box.style.height = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.height[3];\n                            box.style.borderRadius = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.borderRadius.full;\n                            box.style.marginRight = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.margin[2];\n                            box.style.borderWidth = \"3px\";\n                            box.style.borderColor = item.fillStyle;\n                            box.style.pointerEvents = \"none\";\n                            // Label\n                            const labelContainer = document.createElement(\"span\");\n                            labelContainer.style.display = \"flex\";\n                            labelContainer.style.alignItems = \"center\";\n                            const value = document.createElement(\"span\");\n                            value.classList.add(\"text-black\");\n                            value.style.fontSize = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.fontSize[\"3xl\"][0];\n                            value.style.lineHeight = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.fontSize[\"3xl\"][1].lineHeight;\n                            value.style.fontWeight = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.fontWeight.bold;\n                            value.style.marginRight = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.margin[2];\n                            value.style.pointerEvents = \"none\";\n                            const label = document.createElement(\"span\");\n                            label.classList.add(\"text-black\");\n                            label.style.fontSize = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.fontSize.sm[0];\n                            label.style.lineHeight = (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.tailwindConfig)().theme.fontSize.sm[1].lineHeight;\n                            const theValue = c.data.datasets[item.datasetIndex].data.reduce((a, b)=>a + b, 0);\n                            const valueText = document.createTextNode((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_1__.formatValue)(theValue));\n                            const labelText = document.createTextNode(item.text);\n                            value.appendChild(valueText);\n                            label.appendChild(labelText);\n                            li.appendChild(button);\n                            button.appendChild(box);\n                            button.appendChild(labelContainer);\n                            labelContainer.appendChild(value);\n                            labelContainer.appendChild(label);\n                            ul.appendChild(li);\n                        });\n                    }\n                }\n            ]\n        });\n        setChart(newChart);\n        return ()=>newChart.destroy();\n    // eslint-disable-next-line react-hooks/exhaustive-deps\n    }, []);\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)((react__WEBPACK_IMPORTED_MODULE_2___default().Fragment), {\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                className: \"px-5 py-3\",\n                children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"ul\", {\n                    ref: legend,\n                    className: \"flex flex-wrap\"\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\charts\\\\BarChart.tsx\",\n                    lineNumber: 198,\n                    columnNumber: 9\n                }, undefined)\n            }, void 0, false, {\n                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\charts\\\\BarChart.tsx\",\n                lineNumber: 197,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                className: \"grow\",\n                children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"canvas\", {\n                    ref: canvas,\n                    width: width,\n                    height: height\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\charts\\\\BarChart.tsx\",\n                    lineNumber: 201,\n                    columnNumber: 9\n                }, undefined)\n            }, void 0, false, {\n                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\charts\\\\BarChart.tsx\",\n                lineNumber: 200,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true, {\n        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\charts\\\\BarChart.tsx\",\n        lineNumber: 196,\n        columnNumber: 5\n    }, undefined);\n};\n_s(BarChart, \"Zse0wAhIQjOYyy68uWJrxGCM07w=\");\n_c = BarChart;\n/* harmony default export */ __webpack_exports__[\"default\"] = (BarChart);\nvar _c;\n$RefreshReg$(_c, \"BarChart\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevSignature = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevSignature) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports signature on update so we can compare the boundary\n                // signatures. We avoid saving exports themselves since it causes memory leaks (https://github.com/vercel/next.js/pull/53797)\n                module.hot.dispose(function (data) {\n                    data.prevSignature =\n                        self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports);\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevSignature !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevSignature, self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports))) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevSignature !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiKGFwcC1wYWdlcy1icm93c2VyKS8uL2FwcC9jb21wb25lbnRzL2NoYXJ0cy9CYXJDaGFydC50c3giLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7O0FBRWdFO0FBUzlDO0FBQ3lDO0FBQ2I7QUFROUNJLDJDQUFLQSxDQUFDVSxRQUFRLENBQ1paLG1EQUFhQSxFQUNiQyxnREFBVUEsRUFDVkcsaURBQVdBLEVBQ1hDLCtDQUFTQSxFQUNUQyw2Q0FBT0EsRUFDUEgsNENBQU1BO0FBR1IsTUFBTVUsV0FBb0M7UUFBQyxFQUFFQyxJQUFJLEVBQUVDLEtBQUssRUFBRUMsTUFBTSxFQUFFOztJQUNoRSxNQUFNLENBQUNDLE9BQU9DLFNBQVMsR0FBR1IsK0NBQVFBLENBQUM7SUFDbkMsTUFBTVMsU0FBU1YsNkNBQU1BLENBQUM7SUFDdEIsTUFBTVcsU0FBU1gsNkNBQU1BLENBQUM7SUFDdEIsTUFBTSxFQUNKWSxTQUFTLEVBQ1RDLFNBQVMsRUFDVEMsZ0JBQWdCLEVBQ2hCQyxjQUFjLEVBQ2RDLGtCQUFrQixFQUNuQixHQUFHZCx1REFBV0E7SUFFZkgsZ0RBQVNBLENBQUM7UUFDUixNQUFNa0IsTUFBV1AsT0FBT1EsT0FBTztRQUMvQiwwQ0FBMEM7UUFDMUMsTUFBTUMsV0FBZ0IsSUFBSTFCLDJDQUFLQSxDQUFDd0IsS0FBSztZQUNuQ0csTUFBTTtZQUNOZixNQUFNQTtZQUNOZ0IsU0FBUztnQkFDUEMsUUFBUTtvQkFDTkMsU0FBUzt3QkFDUEMsS0FBSzt3QkFDTEMsUUFBUTt3QkFDUkMsTUFBTTt3QkFDTkMsT0FBTztvQkFDVDtnQkFDRjtnQkFDQUMsUUFBUTtvQkFDTkMsR0FBRzt3QkFDREMsUUFBUTs0QkFDTkMsU0FBUzt3QkFDWDt3QkFDQUMsT0FBTzs0QkFDTEMsZUFBZTs0QkFDZkMsVUFBVSxDQUFDQyxRQUFVOUMsNkRBQVdBLENBQUM4Qzs0QkFDakNDLE9BQU94QixVQUFVeUIsSUFBSTt3QkFDdkI7d0JBQ0FDLE1BQU07NEJBQ0pGLE9BQU92QixVQUFVd0IsSUFBSTt3QkFD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QUFDLEVBQUUsQ0FBQ0gsVUFBVTs0QkFDbEQsTUFBTUksV0FBV3hDLEVBQUUxRCxJQUFJLENBQUNtRyxRQUFRLENBQUNoQyxLQUFLYyxZQUFZLENBQUMsQ0FBQ2pGLElBQUksQ0FBQ29HLE1BQU0sQ0FDN0QsQ0FBQ0MsR0FBV0MsSUFBY0QsSUFBSUMsR0FDOUI7NEJBRUYsTUFBTUMsWUFBWWxDLFNBQVNtQyxjQUFjLENBQUN4SCw2REFBV0EsQ0FBQ2tIOzRCQUN0RCxNQUFNTyxZQUFZcEMsU0FBU21DLGNBQWMsQ0FBQ3JDLEtBQUt1QyxJQUFJOzRCQUNuRDVFLE1BQU02RSxXQUFXLENBQUNKOzRCQUNsQjVELE1BQU1nRSxXQUFXLENBQUNGOzRCQUNsQnJDLEdBQUd1QyxXQUFXLENBQUNoQzs0QkFDZkEsT0FBT2dDLFdBQVcsQ0FBQ3ZCOzRCQUNuQlQsT0FBT2dDLFdBQVcsQ0FBQ2pCOzRCQUNuQkEsZUFBZWlCLFdBQVcsQ0FBQzdFOzRCQUMzQjRELGVBQWVpQixXQUFXLENBQUNoRTs0QkFDM0JpQixHQUFHK0MsV0FBVyxDQUFDdkM7d0JBQ2pCO29CQUNGO2dCQUNGO2FBQ0Q7UUFDSDtRQUNBaEUsU0FBU1U7UUFDVCxPQUFPLElBQU1BLFNBQVM4RixPQUFPO0lBQzdCLHVEQUF1RDtJQUN6RCxHQUFHLEVBQUU7SUFFTCxxQkFDRSw4REFBQ25ILHVEQUFjOzswQkFDYiw4REFBQ3FIO2dCQUFJQyxXQUFVOzBCQUNiLDRFQUFDbkQ7b0JBQUdvRCxLQUFLMUc7b0JBQVF5RyxXQUFVOzs7Ozs7Ozs7OzswQkFFN0IsOERBQUNEO2dCQUFJQyxXQUFVOzBCQUNiLDRFQUFDMUc7b0JBQU8yRyxLQUFLM0c7b0JBQVFKLE9BQU9BO29CQUFPQyxRQUFRQTs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFJbkQ7R0E5S01IO0tBQUFBO0FBZ0xOLCtEQUFlQSxRQUFRQSxFQUFDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vX05fRS8uL2FwcC9jb21wb25lbnRzL2NoYXJ0cy9CYXJDaGFydC50c3g/NjMwYiJdLCJzb3VyY2VzQ29udGVudCI6WyJcInVzZSBjbGllbnRcIjtcclxuXHJcbmltcG9ydCB7IGZvcm1hdFZhbHVlLCB0YWlsd2luZENvbmZpZyB9IGZyb20gXCJAL2FwcC91dGlscy9VdGlsc1wiO1xyXG5pbXBvcnQge1xyXG4gIEJhckNvbnRyb2xsZXIsXHJcbiAgQmFyRWxlbWVudCxcclxuICBDaGFydCxcclxuICBMZWdlbmQsXHJcbiAgTGluZWFyU2NhbGUsXHJcbiAgVGltZVNjYWxlLFxyXG4gIFRvb2x0aXAsXHJcbn0gZnJvbSBcImNoYXJ0LmpzXCI7XHJcbmltcG9ydCBSZWFjdCwgeyB1c2VFZmZlY3QsIHVzZVJlZiwgdXNlU3RhdGUgfSBmcm9tIFwicmVhY3RcIjtcclxuaW1wb3J0IHsgY2hhcnRDb2xvcnMgfSBmcm9tIFwiLi9DaGFydGpzQ29uZmlnXCI7XHJcblxyXG5pbnRlcmZhY2UgQmFyQ2hhcnRQcm9wcyB7XHJcbiAgZGF0YTogYW55O1xyXG4gIHdpZHRoOiBudW1iZXI7XHJcbiAgaGVpZ2h0OiBudW1iZXI7XHJcbn1cclxuXHJcbkNoYXJ0LnJlZ2lzdGVyKFxyXG4gIEJhckNvbnRyb2xsZXIsXHJcbiAgQmFyRWxlbWVudCxcclxuICBMaW5lYXJTY2FsZSxcclxuICBUaW1lU2NhbGUsXHJcbiAgVG9vbHRpcCxcclxuICBMZWdlbmRcclxuKTtcclxuXHJcbmNvbnN0IEJhckNoYXJ0OiBSZWFjdC5GQzxCYXJDaGFydFByb3BzPiA9ICh7IGRhdGEsIHdpZHRoLCBoZWlnaHQgfSkgPT4ge1xyXG4gIGNvbnN0IFtjaGFydCwgc2V0Q2hhcnRdID0gdXNlU3RhdGUobnVsbCk7XHJcbiAgY29uc3QgY2FudmFzID0gdXNlUmVmKG51bGwpO1xyXG4gIGNvbnN0IGxlZ2VuZCA9IHVzZVJlZihudWxsKTtcclxuICBjb25zdCB7XHJcbiAgICB0ZXh0Q29sb3IsXHJcbiAgICBncmlkQ29sb3IsXHJcbiAgICB0b29sdGlwQm9keUNvbG9yLFxyXG4gICAgdG9vbHRpcEJnQ29sb3IsXHJcbiAgICB0b29sdGlwQm9yZGVyQ29sb3IsXHJcbiAgfSA9IGNoYXJ0Q29sb3JzO1xyXG5cclxuICB1c2VFZmZlY3QoKCkgPT4ge1xyXG4gICAgY29uc3QgY3R4OiBhbnkgPSBjYW52YXMuY3VycmVudDtcclxuICAgIC8vIGVzbGludC1kaXNhYmxlLW5leHQtbGluZSBuby11bnVzZWQtdmFyc1xyXG4gICAgY29uc3QgbmV3Q2hhcnQ6IGFueSA9IG5ldyBDaGFydChjdHgsIHtcclxuICAgICAgdHlwZTogXCJiYXJcIixcclxuICAgICAgZGF0YTogZGF0YSxcclxuICAgICAgb3B0aW9uczoge1xyXG4gICAgICAgIGxheW91dDoge1xyXG4gICAgICAgICAgcGFkZGluZzoge1xyXG4gICAgICAgICAgICB0b3A6IDEyLFxyXG4gICAgICAgICAgICBib3R0b206IDE2LFxyXG4gICAgICAgICAgICBsZWZ0OiAyMCxcclxuICAgICAgICAgICAgcmlnaHQ6IDIwLFxyXG4gICAgICAgICAgfSxcclxuICAgICAgICB9LFxyXG4gICAgICAgIHNjYWxlczoge1xyXG4gICAgICAgICAgeToge1xyXG4gICAgICAgICAgICBib3JkZXI6IHtcclxuICAgICAgICAgICAgICBkaXNwbGF5OiBmYWxzZSxcclxuICAgICAgICAgICAgfSxcclxuICAgICAgICAgICAgdGlja3M6IHtcclxuICAgICAgICAgICAgICBtYXhUaWNrc0xpbWl0OiA1LFxyXG4gICAgICAgICAgICAgIGNhbGxiYWNrOiAodmFsdWUpID0+IGZvcm1hdFZhbHVlKHZhbHVlKSxcclxuICAgICAgICAgICAgICBjb2xvcjogdGV4dENvbG9yLmRhcmssXHJcbiAgICAgICAgICAgIH0sXHJcbiAgICAgICAgICAgIGdyaWQ6IHtcclxuICAgICAgICAgICAgICBjb2xvcjogZ3JpZENvbG9yLmRhcmssXHJcbiAgICAgICAgICAgIH0sXHJcbiAgICAgICAgICB9LFxyXG4gICAgICAgICAgeDoge1xyXG4gICAgICAgICAgICB0eXBlOiBcInRpbWVcIixcclxuICAgICAgICAgICAgdGltZToge1xyXG4gICAgICAgICAgICAgIHBhcnNlcjogXCJNTS1ERC1ZWVlZXCIsXHJcbiAgICAgICAgICAgICAgdW5pdDogXCJtb250aFwiLFxyXG4gICAgICAgICAgICAgIGRpc3BsYXlGb3JtYXRzOiB7XHJcbiAgICAgICAgICAgICAgICBtb250aDogXCJNTU0gWVlcIixcclxuICAgICAgICAgICAgICB9LFxyXG4gICAgICAgICAgICB9LFxyXG4gICAgICAgICAgICBib3JkZXI6IHtcclxuICAgICAgICAgICAgICBkaXNwbGF5OiBmYWxzZSxcclxuICAgICAgICAgICAgfSxcclxuICAgICAgICAgICAgZ3JpZDoge1xyXG4gICAgICAgICAgICAgIGRpc3BsYXk6IGZhbHNlLFxyXG4gICAgICAgICAgICB9LFxyXG4gICAgICAgICAgICB0aWNrczoge1xyXG4gICAgICAgICAgICAgIGNvbG9yOiB0ZXh0Q29sb3IuZGFyayxcclxuICAgICAgICAgICAgfSxcclxuICAgICAgICAgIH0sXHJcbiAgICAgICAgfSxcclxuICAgICAgICBwbHVnaW5zOiB7XHJcbiAgICAgICAgICBsZWdlbmQ6IHtcclxuICAgICAgICAgICAgZGlzcGxheTogZmFsc2UsXHJcbiAgICAgICAgICB9LFxyXG4gICAgICAgICAgdG9vbHRpcDoge1xyXG4gICAgICAgICAgICBjYWxsYmFja3M6IHtcclxuICAgICAgICAgICAgICAvL3RpdGxlOiAoKSA9PiBmYWxzZSwgLy8gRGlzYWJsZSB0b29sdGlwIHRpdGxlXHJcbiAgICAgICAgICAgICAgbGFiZWw6IChjb250ZXh0KSA9PiBmb3JtYXRWYWx1ZShjb250ZXh0LnBhcnNlZC55KSxcclxuICAgICAgICAgICAgfSxcclxuICAgICAgICAgICAgYm9keUNvbG9yOiB0b29sdGlwQm9keUNvbG9yLmRhcmssXHJcbiAgICAgICAgICAgIGJhY2tncm91bmRDb2xvcjogdG9vbHRpcEJnQ29sb3IuZGFyayxcclxuICAgICAgICAgICAgYm9yZGVyQ29sb3I6IHRvb2x0aXBCb3JkZXJDb2xvci5kYXJrLFxyXG4gICAgICAgICAgfSxcclxuICAgICAgICB9LFxyXG4gICAgICAgIGludGVyYWN0aW9uOiB7XHJcbiAgICAgICAgICBpbnRlcnNlY3Q6IGZhbHNlLFxyXG4gICAgICAgICAgbW9kZTogXCJuZWFyZXN0XCIsXHJcbiAgICAgICAgfSxcclxuICAgICAgICBhbmltYXRpb246IHtcclxuICAgICAgICAgIGR1cmF0aW9uOiA1MDAsXHJcbiAgICAgICAgfSxcclxuICAgICAgICBtYWludGFpbkFzcGVjdFJhdGlvOiBmYWxzZSxcclxuICAgICAgICByZXNpemVEZWxheTogMjAwLFxyXG4gICAgICB9LFxyXG4gICAgICBwbHVnaW5zOiBbXHJcbiAgICAgICAge1xyXG4gICAgICAgICAgaWQ6IFwiaHRtbExlZ2VuZFwiLFxyXG4gICAgICAgICAgYWZ0ZXJVcGRhdGUoYzogYW55LCBhcmdzLCBvcHRpb25zKSB7XHJcbiAgICAgICAgICAgIGNvbnN0IHVsOiBhbnkgPSBsZWdlbmQuY3VycmVudDtcclxuICAgICAgICAgICAgaWYgKCF1bCkgcmV0dXJuO1xyXG4gICAgICAgICAgICAvLyBSZW1vdmUgb2xkIGxlZ2VuZCBpdGVtc1xyXG4gICAgICAgICAgICB3aGlsZSAodWwuZmlyc3RDaGlsZCkge1xyXG4gICAgICAgICAgICAgIHVsLmZpcnN0Q2hpbGQucmVtb3ZlKCk7XHJcbiAgICAgICAgICAgIH1cclxuICAgICAgICAgICAgLy8gUmV1c2UgdGhlIGJ1aWx0LWluIGxlZ2VuZEl0ZW1zIGdlbmVyYXRvclxyXG4gICAgICAgICAgICBjb25zdCBpdGVtczogYW55ID1cclxuICAgICAgICAgICAgICBjLm9wdGlvbnMucGx1Z2lucy5sZWdlbmQubGFiZWxzLmdlbmVyYXRlTGFiZWxzKGMpO1xyXG4gICAgICAgICAgICBpdGVtcy5mb3JFYWNoKChpdGVtOiBhbnkpID0+IHtcclxuICAgICAgICAgICAgICBjb25zdCBsaSA9IGRvY3VtZW50LmNyZWF0ZUVsZW1lbnQoXCJsaVwiKTtcclxuICAgICAgICAgICAgICBsaS5zdHlsZS5tYXJnaW5SaWdodCA9IHRhaWx3aW5kQ29uZmlnKCkudGhlbWUubWFyZ2luWzRdO1xyXG4gICAgICAgICAgICAgIC8vIEJ1dHRvbiBlbGVtZW50XHJcbiAgICAgICAgICAgICAgY29uc3QgYnV0dG9uID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcImJ1dHRvblwiKTtcclxuICAgICAgICAgICAgICBidXR0b24uc3R5bGUuZGlzcGxheSA9IFwiaW5saW5lLWZsZXhcIjtcclxuICAgICAgICAgICAgICBidXR0b24uc3R5bGUuYWxpZ25JdGVtcyA9IFwiY2VudGVyXCI7XHJcbiAgICAgICAgICAgICAgYnV0dG9uLnN0eWxlLm9wYWNpdHkgPSBpdGVtLmhpZGRlbiA/IFwiLjNcIiA6IFwiXCI7XHJcbiAgICAgICAgICAgICAgYnV0dG9uLm9uY2xpY2sgPSAoKSA9PiB7XHJcbiAgICAgICAgICAgICAgICBjLnNldERhdGFzZXRWaXNpYmlsaXR5KFxyXG4gICAgICAgICAgICAgICAgICBpdGVtLmRhdGFzZXRJbmRleCxcclxuICAgICAgICAgICAgICAgICAgIWMuaXNEYXRhc2V0VmlzaWJsZShpdGVtLmRhdGFzZXRJbmRleClcclxuICAgICAgICAgICAgICAgICk7XHJcbiAgICAgICAgICAgICAgICBjLnVwZGF0ZSgpO1xyXG4gICAgICAgICAgICAgIH07XHJcbiAgICAgICAgICAgICAgLy8gQ29sb3IgYm94XHJcbiAgICAgICAgICAgICAgY29uc3QgYm94ID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcInNwYW5cIik7XHJcbiAgICAgICAgICAgICAgYm94LnN0eWxlLmRpc3BsYXkgPSBcImJsb2NrXCI7XHJcbiAgICAgICAgICAgICAgYm94LnN0eWxlLndpZHRoID0gdGFpbHdpbmRDb25maWcoKS50aGVtZS53aWR0aFszXTtcclxuICAgICAgICAgICAgICBib3guc3R5bGUuaGVpZ2h0ID0gdGFpbHdpbmRDb25maWcoKS50aGVtZS5oZWlnaHRbM107XHJcbiAgICAgICAgICAgICAgYm94LnN0eWxlLmJvcmRlclJhZGl1cyA9IHRhaWx3aW5kQ29uZmlnKCkudGhlbWUuYm9yZGVyUmFkaXVzLmZ1bGw7XHJcbiAgICAgICAgICAgICAgYm94LnN0eWxlLm1hcmdpblJpZ2h0ID0gdGFpbHdpbmRDb25maWcoKS50aGVtZS5tYXJnaW5bMl07XHJcbiAgICAgICAgICAgICAgYm94LnN0eWxlLmJvcmRlcldpZHRoID0gXCIzcHhcIjtcclxuICAgICAgICAgICAgICBib3guc3R5bGUuYm9yZGVyQ29sb3IgPSBpdGVtLmZpbGxTdHlsZTtcclxuICAgICAgICAgICAgICBib3guc3R5bGUucG9pbnRlckV2ZW50cyA9IFwibm9uZVwiO1xyXG4gICAgICAgICAgICAgIC8vIExhYmVsXHJcbiAgICAgICAgICAgICAgY29uc3QgbGFiZWxDb250YWluZXIgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwic3BhblwiKTtcclxuICAgICAgICAgICAgICBsYWJlbENvbnRhaW5lci5zdHlsZS5kaXNwbGF5ID0gXCJmbGV4XCI7XHJcbiAgICAgICAgICAgICAgbGFiZWxDb250YWluZXIuc3R5bGUuYWxpZ25JdGVtcyA9IFwiY2VudGVyXCI7XHJcbiAgICAgICAgICAgICAgY29uc3QgdmFsdWUgPSBkb2N1bWVudC5jcmVhdGVFbGVtZW50KFwic3BhblwiKTtcclxuICAgICAgICAgICAgICB2YWx1ZS5jbGFzc0xpc3QuYWRkKFwidGV4dC1ibGFja1wiKTtcclxuICAgICAgICAgICAgICB2YWx1ZS5zdHlsZS5mb250U2l6ZSA9IHRhaWx3aW5kQ29uZmlnKCkudGhlbWUuZm9udFNpemVbXCIzeGxcIl1bMF07XHJcbiAgICAgICAgICAgICAgdmFsdWUuc3R5bGUubGluZUhlaWdodCA9XHJcbiAgICAgICAgICAgICAgICB0YWlsd2luZENvbmZpZygpLnRoZW1lLmZvbnRTaXplW1wiM3hsXCJdWzFdLmxpbmVIZWlnaHQ7XHJcbiAgICAgICAgICAgICAgdmFsdWUuc3R5bGUuZm9udFdlaWdodCA9IHRhaWx3aW5kQ29uZmlnKCkudGhlbWUuZm9udFdlaWdodC5ib2xkO1xyXG4gICAgICAgICAgICAgIHZhbHVlLnN0eWxlLm1hcmdpblJpZ2h0ID0gdGFpbHdpbmRDb25maWcoKS50aGVtZS5tYXJnaW5bMl07XHJcbiAgICAgICAgICAgICAgdmFsdWUuc3R5bGUucG9pbnRlckV2ZW50cyA9IFwibm9uZVwiO1xyXG4gICAgICAgICAgICAgIGNvbnN0IGxhYmVsID0gZG9jdW1lbnQuY3JlYXRlRWxlbWVudChcInNwYW5cIik7XHJcbiAgICAgICAgICAgICAgbGFiZWwuY2xhc3NMaXN0LmFkZChcInRleHQtYmxhY2tcIik7XHJcbiAgICAgICAgICAgICAgbGFiZWwuc3R5bGUuZm9udFNpemUgPSB0YWlsd2luZENvbmZpZygpLnRoZW1lLmZvbnRTaXplLnNtWzBdO1xyXG4gICAgICAgICAgICAgIGxhYmVsLnN0eWxlLmxpbmVIZWlnaHQgPVxyXG4gICAgICAgICAgICAgICAgdGFpbHdpbmRDb25maWcoKS50aGVtZS5mb250U2l6ZS5zbVsxXS5saW5lSGVpZ2h0O1xyXG4gICAgICAgICAgICAgIGNvbnN0IHRoZVZhbHVlID0gYy5kYXRhLmRhdGFzZXRzW2l0ZW0uZGF0YXNldEluZGV4XS5kYXRhLnJlZHVjZShcclxuICAgICAgICAgICAgICAgIChhOiBudW1iZXIsIGI6IG51bWJlcikgPT4gYSArIGIsXHJcbiAgICAgICAgICAgICAgICAwXHJcbiAgICAgICAgICAgICAgKTtcclxuICAgICAgICAgICAgICBjb25zdCB2YWx1ZVRleHQgPSBkb2N1bWVudC5jcmVhdGVUZXh0Tm9kZShmb3JtYXRWYWx1ZSh0aGVWYWx1ZSkpO1xyXG4gICAgICAgICAgICAgIGNvbnN0IGxhYmVsVGV4dCA9IGRvY3VtZW50LmNyZWF0ZVRleHROb2RlKGl0ZW0udGV4dCk7XHJcbiAgICAgICAgICAgICAgdmFsdWUuYXBwZW5kQ2hpbGQodmFsdWVUZXh0KTtcclxuICAgICAgICAgICAgICBsYWJlbC5hcHBlbmRDaGlsZChsYWJlbFRleHQpO1xyXG4gICAgICAgICAgICAgIGxpLmFwcGVuZENoaWxkKGJ1dHRvbik7XHJcbiAgICAgICAgICAgICAgYnV0dG9uLmFwcGVuZENoaWxkKGJveCk7XHJcbiAgICAgICAgICAgICAgYnV0dG9uLmFwcGVuZENoaWxkKGxhYmVsQ29udGFpbmVyKTtcclxuICAgICAgICAgICAgICBsYWJlbENvbnRhaW5lci5hcHBlbmRDaGlsZCh2YWx1ZSk7XHJcbiAgICAgICAgICAgICAgbGFiZWxDb250YWluZXIuYXBwZW5kQ2hpbGQobGFiZWwpO1xyXG4gICAgICAgICAgICAgIHVsLmFwcGVuZENoaWxkKGxpKTtcclxuICAgICAgICAgICAgfSk7XHJcbiAgICAgICAgICB9LFxyXG4gICAgICAgIH0sXHJcbiAgICAgIF0sXHJcbiAgICB9KTtcclxuICAgIHNldENoYXJ0KG5ld0NoYXJ0KTtcclxuICAgIHJldHVybiAoKSA9PiBuZXdDaGFydC5kZXN0cm95KCk7XHJcbiAgICAvLyBlc2xpbnQtZGlzYWJsZS1uZXh0LWxpbmUgcmVhY3QtaG9va3MvZXhoYXVzdGl2ZS1kZXBzXHJcbiAgfSwgW10pO1xyXG5cclxuICByZXR1cm4gKFxyXG4gICAgPFJlYWN0LkZyYWdtZW50PlxyXG4gICAgICA8ZGl2IGNsYXNzTmFtZT1cInB4LTUgcHktM1wiPlxyXG4gICAgICAgIDx1bCByZWY9e2xlZ2VuZH0gY2xhc3NOYW1lPVwiZmxleCBmbGV4LXdyYXBcIj48L3VsPlxyXG4gICAgICA8L2Rpdj5cclxuICAgICAgPGRpdiBjbGFzc05hbWU9XCJncm93XCI+XHJcbiAgICAgICAgPGNhbnZhcyByZWY9e2NhbnZhc30gd2lkdGg9e3dpZHRofSBoZWlnaHQ9e2hlaWdodH0+PC9jYW52YXM+XHJcbiAgICAgIDwvZGl2PlxyXG4gICAgPC9SZWFjdC5GcmFnbWVudD5cclxuICApO1xyXG59O1xyXG5cclxuZXhwb3J0IGRlZmF1bHQgQmFyQ2hhcnQ7XHJcbiJdLCJuYW1lcyI6WyJmb3JtYXRWYWx1ZSIsInRhaWx3aW5kQ29uZmlnIiwiQmFyQ29udHJvbGxlciIsIkJhckVsZW1lbnQiLCJDaGFydCIsIkxlZ2VuZCIsIkxpbmVhclNjYWxlIiwiVGltZVNjYWxlIiwiVG9vbHRpcCIsIlJlYWN0IiwidXNlRWZmZWN0IiwidXNlUmVmIiwidXNlU3RhdGUiLCJjaGFydENvbG9ycyIsInJlZ2lzdGVyIiwiQmFyQ2hhcnQiLCJkYXRhIiwid2lkdGgiLCJoZWlnaHQiLCJjaGFydCIsInNldENoYXJ0IiwiY2FudmFzIiwibGVnZW5kIiwidGV4dENvbG9yIiwiZ3JpZENvbG9yIiwidG9vbHRpcEJvZHlDb2xvciIsInRvb2x0aXBCZ0NvbG9yIiwidG9vbHRpcEJvcmRlckNvbG9yIiwiY3R4IiwiY3VycmVudCIsIm5ld0NoYXJ0IiwidHlwZSIsIm9wdGlvbnMiLCJsYXlvdXQiLCJwYWRkaW5nIiwidG9wIiwiYm90dG9tIiwibGVmdCIsInJpZ2h0Iiwic2NhbGVzIiwieSIsImJvcmRlciIsImRpc3BsYXkiLCJ0aWNrcyIsIm1heFRpY2tzTGltaXQiLCJjYWxsYmFjayIsInZhbHVlIiwiY29sb3IiLCJkYXJrIiwiZ3JpZCIsIngiLCJ0aW1lIiwicGFyc2VyIiwidW5pdCIsImRpc3BsYXlGb3JtYXRzIiwibW9udGgiLCJwbHVnaW5zIiwidG9vbHRpcCIsImNhbGxiYWNrcyIsImxhYmVsIiwiY29udGV4dCIsInBhcnNlZCIsImJvZHlDb2xvciIsImJhY2tncm91bmRDb2xvciIsImJvcmRlckNvbG9yIiwiaW50ZXJhY3Rpb24iLCJpbnRlcnNlY3QiLCJtb2RlIiwiYW5pbWF0aW9uIiwiZHVyYXRpb24iLCJtYWludGFpbkFzcGVjdFJhdGlvIiwicmVzaXplRGVsYXkiLCJpZCIsImFmdGVyVXBkYXRlIiwiYyIsImFyZ3MiLCJ1bCIsImZpcnN0Q2hpbGQiLCJyZW1vdmUiLCJpdGVtcyIsImxhYmVscyIsImdlbmVyYXRlTGFiZWxzIiwiZm9yRWFjaCIsIml0ZW0iLCJsaSIsImRvY3VtZW50IiwiY3JlYXRlRWxlbWVudCIsInN0eWxlIiwibWFyZ2luUmlnaHQiLCJ0aGVtZSIsIm1hcmdpbiIsImJ1dHRvbiIsImFsaWduSXRlbXMiLCJvcGFjaXR5IiwiaGlkZGVuIiwib25jbGljayIsInNldERhdGFzZXRWaXNpYmlsaXR5IiwiZGF0YXNldEluZGV4IiwiaXNEYXRhc2V0VmlzaWJsZSIsInVwZGF0ZSIsImJveCIsImJvcmRlclJhZGl1cyIsImZ1bGwiLCJib3JkZXJXaWR0aCIsImZpbGxTdHlsZSIsInBvaW50ZXJFdmVudHMiLCJsYWJlbENvbnRhaW5lciIsImNsYXNzTGlzdCIsImFkZCIsImZvbnRTaXplIiwibGluZUhlaWdodCIsImZvbnRXZWlnaHQiLCJib2xkIiwic20iLCJ0aGVWYWx1ZSIsImRhdGFzZXRzIiwicmVkdWNlIiwiYSIsImIiLCJ2YWx1ZVRleHQiLCJjcmVhdGVUZXh0Tm9kZSIsImxhYmVsVGV4dCIsInRleHQiLCJhcHBlbmRDaGlsZCIsImRlc3Ryb3kiLCJGcmFnbWVudCIsImRpdiIsImNsYXNzTmFtZSIsInJlZiJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///(app-pages-browser)/./app/components/charts/BarChart.tsx\n"));

/***/ })

});
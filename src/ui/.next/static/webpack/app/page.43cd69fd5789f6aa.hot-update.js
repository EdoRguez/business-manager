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

/***/ "(app-pages-browser)/./app/components/cards/SimpleLineChartCard.tsx":
/*!******************************************************!*\
  !*** ./app/components/cards/SimpleLineChartCard.tsx ***!
  \******************************************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"(app-pages-browser)/./node_modules/next/dist/compiled/react/jsx-dev-runtime.js\");\n/* harmony import */ var _charts_LineChartSimple__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! ../charts/LineChartSimple */ \"(app-pages-browser)/./app/components/charts/LineChartSimple.tsx\");\n/* harmony import */ var _app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @/app/utils/Utils */ \"(app-pages-browser)/./app/utils/Utils.ts\");\n/* __next_internal_client_entry_do_not_use__ default auto */ \n\n\nconst SimpleLineChartCard = ()=>{\n    const chartData = {\n        labels: [\n            \"12-01-2020\",\n            \"01-01-2021\",\n            \"02-01-2021\",\n            \"03-01-2021\",\n            \"04-01-2021\",\n            \"05-01-2021\",\n            \"06-01-2021\",\n            \"07-01-2021\",\n            \"08-01-2021\",\n            \"09-01-2021\",\n            \"10-01-2021\",\n            \"11-01-2021\",\n            \"12-01-2021\",\n            \"01-01-2022\",\n            \"02-01-2022\",\n            \"03-01-2022\",\n            \"04-01-2022\",\n            \"05-01-2022\",\n            \"06-01-2022\",\n            \"07-01-2022\",\n            \"08-01-2022\",\n            \"09-01-2022\",\n            \"10-01-2022\",\n            \"11-01-2022\",\n            \"12-01-2022\",\n            \"01-01-2023\"\n        ],\n        datasets: [\n            // Indigo line\n            {\n                data: [\n                    732,\n                    610,\n                    610,\n                    504,\n                    504,\n                    504,\n                    349,\n                    349,\n                    504,\n                    342,\n                    504,\n                    610,\n                    391,\n                    192,\n                    154,\n                    273,\n                    191,\n                    191,\n                    126,\n                    263,\n                    349,\n                    252,\n                    423,\n                    622,\n                    470,\n                    532\n                ],\n                fill: true,\n                backgroundColor: \"rgba(\".concat((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.hexToRGB)((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.tailwindConfig)().theme.colors.blue[500]), \", 0.08)\"),\n                borderColor: (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.tailwindConfig)().theme.colors.indigo[500],\n                borderWidth: 2,\n                tension: 0,\n                pointRadius: 0,\n                pointHoverRadius: 3,\n                pointBackgroundColor: (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.tailwindConfig)().theme.colors.indigo[500],\n                pointHoverBackgroundColor: (0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.tailwindConfig)().theme.colors.indigo[500],\n                pointBorderWidth: 0,\n                pointHoverBorderWidth: 0,\n                clip: 20\n            },\n            // Gray line\n            {\n                data: [\n                    532,\n                    532,\n                    532,\n                    404,\n                    404,\n                    314,\n                    314,\n                    314,\n                    314,\n                    314,\n                    234,\n                    314,\n                    234,\n                    234,\n                    314,\n                    314,\n                    314,\n                    388,\n                    314,\n                    202,\n                    202,\n                    202,\n                    202,\n                    314,\n                    720,\n                    642\n                ],\n                borderColor: \"rgba(\".concat((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.hexToRGB)((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.tailwindConfig)().theme.colors.slate[500]), \", 0.25)\"),\n                borderWidth: 2,\n                tension: 0,\n                pointRadius: 0,\n                pointHoverRadius: 3,\n                pointBackgroundColor: \"rgba(\".concat((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.hexToRGB)((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.tailwindConfig)().theme.colors.slate[500]), \", 0.25)\"),\n                pointHoverBackgroundColor: \"rgba(\".concat((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.hexToRGB)((0,_app_utils_Utils__WEBPACK_IMPORTED_MODULE_2__.tailwindConfig)().theme.colors.slate[500]), \", 0.25)\"),\n                pointBorderWidth: 0,\n                pointHoverBorderWidth: 0,\n                clip: 20\n            }\n        ]\n    };\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n        className: \"bg-white shadow-lg rounded-md\",\n        children: [\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                className: \"px-5 pt-5\",\n                children: [\n                    /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"h2\", {\n                        className: \"text-md font-semibold text-maincolor mb-2\",\n                        children: \"Acme Plus\"\n                    }, void 0, false, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n                        lineNumber: 87,\n                        columnNumber: 9\n                    }, undefined),\n                    /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                        className: \"text-xs font-semibold text-black uppercase mb-1\",\n                        children: \"Sales\"\n                    }, void 0, false, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n                        lineNumber: 90,\n                        columnNumber: 9\n                    }, undefined),\n                    /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                        className: \"flex items-start\",\n                        children: [\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                                className: \"text-2xl font-bold text-black mr-2\",\n                                children: \"$24,780\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n                                lineNumber: 94,\n                                columnNumber: 11\n                            }, undefined),\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                                className: \"text-sm font-semibold text-white px-2 py-1 bg-emerald-500 rounded-full\",\n                                children: \"+49%\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n                                lineNumber: 97,\n                                columnNumber: 11\n                            }, undefined)\n                        ]\n                    }, void 0, true, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n                        lineNumber: 93,\n                        columnNumber: 9\n                    }, undefined)\n                ]\n            }, void 0, true, {\n                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n                lineNumber: 86,\n                columnNumber: 7\n            }, undefined),\n            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                className: \"grow max-sm:max-h-[128px] xl:max-h-[128px]\",\n                children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_charts_LineChartSimple__WEBPACK_IMPORTED_MODULE_1__[\"default\"], {\n                    data: chartData,\n                    width: 389,\n                    height: 128\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n                    lineNumber: 105,\n                    columnNumber: 9\n                }, undefined)\n            }, void 0, false, {\n                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n                lineNumber: 103,\n                columnNumber: 7\n            }, undefined)\n        ]\n    }, void 0, true, {\n        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\cards\\\\SimpleLineChartCard.tsx\",\n        lineNumber: 85,\n        columnNumber: 5\n    }, undefined);\n};\n_c = SimpleLineChartCard;\n/* harmony default export */ __webpack_exports__[\"default\"] = (SimpleLineChartCard);\nvar _c;\n$RefreshReg$(_c, \"SimpleLineChartCard\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevSignature = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevSignature) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports signature on update so we can compare the boundary\n                // signatures. We avoid saving exports themselves since it causes memory leaks (https://github.com/vercel/next.js/pull/53797)\n                module.hot.dispose(function (data) {\n                    data.prevSignature =\n                        self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports);\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevSignature !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevSignature, self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports))) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevSignature !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiKGFwcC1wYWdlcy1icm93c2VyKS8uL2FwcC9jb21wb25lbnRzL2NhcmRzL1NpbXBsZUxpbmVDaGFydENhcmQudHN4IiwibWFwcGluZ3MiOiI7Ozs7O0FBRXdEO0FBQ0s7QUFFN0QsTUFBTUcsc0JBQXNCO0lBQzFCLE1BQU1DLFlBQVk7UUFDaEJDLFFBQVE7WUFDTjtZQUNBO1lBQ0E7WUFDQTtZQUNBO1lBQ0E7WUFDQTtZQUNBO1lBQ0E7WUFDQTtZQUNBO1lBQ0E7WUFDQTtZQUNBO1lBQ0E7WUFDQTtZQUNBO1lBQ0E7WUFDQTtZQUNBO1lBQ0E7WUFDQTtZQUNBO1lBQ0E7WUFDQTtZQUNBO1NBQ0Q7UUFDREMsVUFBVTtZQUNSLGNBQWM7WUFDZDtnQkFDRUMsTUFBTTtvQkFDSjtvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFDakU7b0JBQUs7b0JBQUs7b0JBQUs7b0JBQUs7b0JBQUs7b0JBQUs7b0JBQUs7b0JBQUs7b0JBQUs7b0JBQUs7b0JBQUs7aUJBQ3hEO2dCQUNEQyxNQUFNO2dCQUNOQyxpQkFBaUIsUUFFZixPQUZ1QlIsMERBQVFBLENBQy9CQyxnRUFBY0EsR0FBR1EsS0FBSyxDQUFDQyxNQUFNLENBQUNDLElBQUksQ0FBQyxJQUFJLEdBQ3ZDO2dCQUNGQyxhQUFhWCxnRUFBY0EsR0FBR1EsS0FBSyxDQUFDQyxNQUFNLENBQUNHLE1BQU0sQ0FBQyxJQUFJO2dCQUN0REMsYUFBYTtnQkFDYkMsU0FBUztnQkFDVEMsYUFBYTtnQkFDYkMsa0JBQWtCO2dCQUNsQkMsc0JBQXNCakIsZ0VBQWNBLEdBQUdRLEtBQUssQ0FBQ0MsTUFBTSxDQUFDRyxNQUFNLENBQUMsSUFBSTtnQkFDL0RNLDJCQUEyQmxCLGdFQUFjQSxHQUFHUSxLQUFLLENBQUNDLE1BQU0sQ0FBQ0csTUFBTSxDQUFDLElBQUk7Z0JBQ3BFTyxrQkFBa0I7Z0JBQ2xCQyx1QkFBdUI7Z0JBQ3ZCQyxNQUFNO1lBQ1I7WUFDQSxZQUFZO1lBQ1o7Z0JBQ0VoQixNQUFNO29CQUNKO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUFLO29CQUNqRTtvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztvQkFBSztpQkFDeEQ7Z0JBQ0RNLGFBQWEsUUFFWCxPQUZtQlosMERBQVFBLENBQzNCQyxnRUFBY0EsR0FBR1EsS0FBSyxDQUFDQyxNQUFNLENBQUNhLEtBQUssQ0FBQyxJQUFJLEdBQ3hDO2dCQUNGVCxhQUFhO2dCQUNiQyxTQUFTO2dCQUNUQyxhQUFhO2dCQUNiQyxrQkFBa0I7Z0JBQ2xCQyxzQkFBc0IsUUFFcEIsT0FGNEJsQiwwREFBUUEsQ0FDcENDLGdFQUFjQSxHQUFHUSxLQUFLLENBQUNDLE1BQU0sQ0FBQ2EsS0FBSyxDQUFDLElBQUksR0FDeEM7Z0JBQ0ZKLDJCQUEyQixRQUV6QixPQUZpQ25CLDBEQUFRQSxDQUN6Q0MsZ0VBQWNBLEdBQUdRLEtBQUssQ0FBQ0MsTUFBTSxDQUFDYSxLQUFLLENBQUMsSUFBSSxHQUN4QztnQkFDRkgsa0JBQWtCO2dCQUNsQkMsdUJBQXVCO2dCQUN2QkMsTUFBTTtZQUNSO1NBQ0Q7SUFDSDtJQUVBLHFCQUNFLDhEQUFDRTtRQUFJQyxXQUFVOzswQkFDYiw4REFBQ0Q7Z0JBQUlDLFdBQVU7O2tDQUNiLDhEQUFDQzt3QkFBR0QsV0FBVTtrQ0FBNEM7Ozs7OztrQ0FHMUQsOERBQUNEO3dCQUFJQyxXQUFVO2tDQUFrRDs7Ozs7O2tDQUdqRSw4REFBQ0Q7d0JBQUlDLFdBQVU7OzBDQUNiLDhEQUFDRDtnQ0FBSUMsV0FBVTswQ0FBcUM7Ozs7OzswQ0FHcEQsOERBQUNEO2dDQUFJQyxXQUFVOzBDQUF5RTs7Ozs7Ozs7Ozs7Ozs7Ozs7OzBCQU01Riw4REFBQ0Q7Z0JBQUlDLFdBQVU7MEJBRWIsNEVBQUMxQiwrREFBZUE7b0JBQUNPLE1BQU1IO29CQUFXd0IsT0FBTztvQkFBS0MsUUFBUTs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFJOUQ7S0F2R00xQjtBQXlHTiwrREFBZUEsbUJBQW1CQSxFQUFDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vX05fRS8uL2FwcC9jb21wb25lbnRzL2NhcmRzL1NpbXBsZUxpbmVDaGFydENhcmQudHN4P2Y0MjAiXSwic291cmNlc0NvbnRlbnQiOlsiXCJ1c2UgY2xpZW50XCI7XHJcblxyXG5pbXBvcnQgTGluZUNoYXJ0U2ltcGxlIGZyb20gXCIuLi9jaGFydHMvTGluZUNoYXJ0U2ltcGxlXCI7XHJcbmltcG9ydCB7IGhleFRvUkdCLCB0YWlsd2luZENvbmZpZyB9IGZyb20gXCJAL2FwcC91dGlscy9VdGlsc1wiO1xyXG5cclxuY29uc3QgU2ltcGxlTGluZUNoYXJ0Q2FyZCA9ICgpID0+IHtcclxuICBjb25zdCBjaGFydERhdGEgPSB7XHJcbiAgICBsYWJlbHM6IFtcclxuICAgICAgXCIxMi0wMS0yMDIwXCIsXHJcbiAgICAgIFwiMDEtMDEtMjAyMVwiLFxyXG4gICAgICBcIjAyLTAxLTIwMjFcIixcclxuICAgICAgXCIwMy0wMS0yMDIxXCIsXHJcbiAgICAgIFwiMDQtMDEtMjAyMVwiLFxyXG4gICAgICBcIjA1LTAxLTIwMjFcIixcclxuICAgICAgXCIwNi0wMS0yMDIxXCIsXHJcbiAgICAgIFwiMDctMDEtMjAyMVwiLFxyXG4gICAgICBcIjA4LTAxLTIwMjFcIixcclxuICAgICAgXCIwOS0wMS0yMDIxXCIsXHJcbiAgICAgIFwiMTAtMDEtMjAyMVwiLFxyXG4gICAgICBcIjExLTAxLTIwMjFcIixcclxuICAgICAgXCIxMi0wMS0yMDIxXCIsXHJcbiAgICAgIFwiMDEtMDEtMjAyMlwiLFxyXG4gICAgICBcIjAyLTAxLTIwMjJcIixcclxuICAgICAgXCIwMy0wMS0yMDIyXCIsXHJcbiAgICAgIFwiMDQtMDEtMjAyMlwiLFxyXG4gICAgICBcIjA1LTAxLTIwMjJcIixcclxuICAgICAgXCIwNi0wMS0yMDIyXCIsXHJcbiAgICAgIFwiMDctMDEtMjAyMlwiLFxyXG4gICAgICBcIjA4LTAxLTIwMjJcIixcclxuICAgICAgXCIwOS0wMS0yMDIyXCIsXHJcbiAgICAgIFwiMTAtMDEtMjAyMlwiLFxyXG4gICAgICBcIjExLTAxLTIwMjJcIixcclxuICAgICAgXCIxMi0wMS0yMDIyXCIsXHJcbiAgICAgIFwiMDEtMDEtMjAyM1wiLFxyXG4gICAgXSxcclxuICAgIGRhdGFzZXRzOiBbXHJcbiAgICAgIC8vIEluZGlnbyBsaW5lXHJcbiAgICAgIHtcclxuICAgICAgICBkYXRhOiBbXHJcbiAgICAgICAgICA3MzIsIDYxMCwgNjEwLCA1MDQsIDUwNCwgNTA0LCAzNDksIDM0OSwgNTA0LCAzNDIsIDUwNCwgNjEwLCAzOTEsIDE5MixcclxuICAgICAgICAgIDE1NCwgMjczLCAxOTEsIDE5MSwgMTI2LCAyNjMsIDM0OSwgMjUyLCA0MjMsIDYyMiwgNDcwLCA1MzIsXHJcbiAgICAgICAgXSxcclxuICAgICAgICBmaWxsOiB0cnVlLFxyXG4gICAgICAgIGJhY2tncm91bmRDb2xvcjogYHJnYmEoJHtoZXhUb1JHQihcclxuICAgICAgICAgIHRhaWx3aW5kQ29uZmlnKCkudGhlbWUuY29sb3JzLmJsdWVbNTAwXVxyXG4gICAgICAgICl9LCAwLjA4KWAsXHJcbiAgICAgICAgYm9yZGVyQ29sb3I6IHRhaWx3aW5kQ29uZmlnKCkudGhlbWUuY29sb3JzLmluZGlnb1s1MDBdLFxyXG4gICAgICAgIGJvcmRlcldpZHRoOiAyLFxyXG4gICAgICAgIHRlbnNpb246IDAsXHJcbiAgICAgICAgcG9pbnRSYWRpdXM6IDAsXHJcbiAgICAgICAgcG9pbnRIb3ZlclJhZGl1czogMyxcclxuICAgICAgICBwb2ludEJhY2tncm91bmRDb2xvcjogdGFpbHdpbmRDb25maWcoKS50aGVtZS5jb2xvcnMuaW5kaWdvWzUwMF0sXHJcbiAgICAgICAgcG9pbnRIb3ZlckJhY2tncm91bmRDb2xvcjogdGFpbHdpbmRDb25maWcoKS50aGVtZS5jb2xvcnMuaW5kaWdvWzUwMF0sXHJcbiAgICAgICAgcG9pbnRCb3JkZXJXaWR0aDogMCxcclxuICAgICAgICBwb2ludEhvdmVyQm9yZGVyV2lkdGg6IDAsXHJcbiAgICAgICAgY2xpcDogMjAsXHJcbiAgICAgIH0sXHJcbiAgICAgIC8vIEdyYXkgbGluZVxyXG4gICAgICB7XHJcbiAgICAgICAgZGF0YTogW1xyXG4gICAgICAgICAgNTMyLCA1MzIsIDUzMiwgNDA0LCA0MDQsIDMxNCwgMzE0LCAzMTQsIDMxNCwgMzE0LCAyMzQsIDMxNCwgMjM0LCAyMzQsXHJcbiAgICAgICAgICAzMTQsIDMxNCwgMzE0LCAzODgsIDMxNCwgMjAyLCAyMDIsIDIwMiwgMjAyLCAzMTQsIDcyMCwgNjQyLFxyXG4gICAgICAgIF0sXHJcbiAgICAgICAgYm9yZGVyQ29sb3I6IGByZ2JhKCR7aGV4VG9SR0IoXHJcbiAgICAgICAgICB0YWlsd2luZENvbmZpZygpLnRoZW1lLmNvbG9ycy5zbGF0ZVs1MDBdXHJcbiAgICAgICAgKX0sIDAuMjUpYCxcclxuICAgICAgICBib3JkZXJXaWR0aDogMixcclxuICAgICAgICB0ZW5zaW9uOiAwLFxyXG4gICAgICAgIHBvaW50UmFkaXVzOiAwLFxyXG4gICAgICAgIHBvaW50SG92ZXJSYWRpdXM6IDMsXHJcbiAgICAgICAgcG9pbnRCYWNrZ3JvdW5kQ29sb3I6IGByZ2JhKCR7aGV4VG9SR0IoXHJcbiAgICAgICAgICB0YWlsd2luZENvbmZpZygpLnRoZW1lLmNvbG9ycy5zbGF0ZVs1MDBdXHJcbiAgICAgICAgKX0sIDAuMjUpYCxcclxuICAgICAgICBwb2ludEhvdmVyQmFja2dyb3VuZENvbG9yOiBgcmdiYSgke2hleFRvUkdCKFxyXG4gICAgICAgICAgdGFpbHdpbmRDb25maWcoKS50aGVtZS5jb2xvcnMuc2xhdGVbNTAwXVxyXG4gICAgICAgICl9LCAwLjI1KWAsXHJcbiAgICAgICAgcG9pbnRCb3JkZXJXaWR0aDogMCxcclxuICAgICAgICBwb2ludEhvdmVyQm9yZGVyV2lkdGg6IDAsXHJcbiAgICAgICAgY2xpcDogMjAsXHJcbiAgICAgIH0sXHJcbiAgICBdLFxyXG4gIH07XHJcblxyXG4gIHJldHVybiAoXHJcbiAgICA8ZGl2IGNsYXNzTmFtZT1cImJnLXdoaXRlIHNoYWRvdy1sZyByb3VuZGVkLW1kXCI+XHJcbiAgICAgIDxkaXYgY2xhc3NOYW1lPVwicHgtNSBwdC01XCI+XHJcbiAgICAgICAgPGgyIGNsYXNzTmFtZT1cInRleHQtbWQgZm9udC1zZW1pYm9sZCB0ZXh0LW1haW5jb2xvciBtYi0yXCI+XHJcbiAgICAgICAgICBBY21lIFBsdXNcclxuICAgICAgICA8L2gyPlxyXG4gICAgICAgIDxkaXYgY2xhc3NOYW1lPVwidGV4dC14cyBmb250LXNlbWlib2xkIHRleHQtYmxhY2sgdXBwZXJjYXNlIG1iLTFcIj5cclxuICAgICAgICAgIFNhbGVzXHJcbiAgICAgICAgPC9kaXY+XHJcbiAgICAgICAgPGRpdiBjbGFzc05hbWU9XCJmbGV4IGl0ZW1zLXN0YXJ0XCI+XHJcbiAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cInRleHQtMnhsIGZvbnQtYm9sZCB0ZXh0LWJsYWNrIG1yLTJcIj5cclxuICAgICAgICAgICAgJDI0LDc4MFxyXG4gICAgICAgICAgPC9kaXY+XHJcbiAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT1cInRleHQtc20gZm9udC1zZW1pYm9sZCB0ZXh0LXdoaXRlIHB4LTIgcHktMSBiZy1lbWVyYWxkLTUwMCByb3VuZGVkLWZ1bGxcIj5cclxuICAgICAgICAgICAgKzQ5JVxyXG4gICAgICAgICAgPC9kaXY+XHJcbiAgICAgICAgPC9kaXY+XHJcbiAgICAgIDwvZGl2PlxyXG4gICAgICB7LyogQ2hhcnQgYnVpbHQgd2l0aCBDaGFydC5qcyAzICovfVxyXG4gICAgICA8ZGl2IGNsYXNzTmFtZT1cImdyb3cgbWF4LXNtOm1heC1oLVsxMjhweF0geGw6bWF4LWgtWzEyOHB4XVwiPlxyXG4gICAgICAgIHsvKiBDaGFuZ2UgdGhlIGhlaWdodCBhdHRyaWJ1dGUgdG8gYWRqdXN0IHRoZSBjaGFydCBoZWlnaHQgKi99XHJcbiAgICAgICAgPExpbmVDaGFydFNpbXBsZSBkYXRhPXtjaGFydERhdGF9IHdpZHRoPXszODl9IGhlaWdodD17MTI4fSAvPlxyXG4gICAgICA8L2Rpdj5cclxuICAgIDwvZGl2PlxyXG4gICk7XHJcbn07XHJcblxyXG5leHBvcnQgZGVmYXVsdCBTaW1wbGVMaW5lQ2hhcnRDYXJkO1xyXG4iXSwibmFtZXMiOlsiTGluZUNoYXJ0U2ltcGxlIiwiaGV4VG9SR0IiLCJ0YWlsd2luZENvbmZpZyIsIlNpbXBsZUxpbmVDaGFydENhcmQiLCJjaGFydERhdGEiLCJsYWJlbHMiLCJkYXRhc2V0cyIsImRhdGEiLCJmaWxsIiwiYmFja2dyb3VuZENvbG9yIiwidGhlbWUiLCJjb2xvcnMiLCJibHVlIiwiYm9yZGVyQ29sb3IiLCJpbmRpZ28iLCJib3JkZXJXaWR0aCIsInRlbnNpb24iLCJwb2ludFJhZGl1cyIsInBvaW50SG92ZXJSYWRpdXMiLCJwb2ludEJhY2tncm91bmRDb2xvciIsInBvaW50SG92ZXJCYWNrZ3JvdW5kQ29sb3IiLCJwb2ludEJvcmRlcldpZHRoIiwicG9pbnRIb3ZlckJvcmRlcldpZHRoIiwiY2xpcCIsInNsYXRlIiwiZGl2IiwiY2xhc3NOYW1lIiwiaDIiLCJ3aWR0aCIsImhlaWdodCJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///(app-pages-browser)/./app/components/cards/SimpleLineChartCard.tsx\n"));

/***/ })

});
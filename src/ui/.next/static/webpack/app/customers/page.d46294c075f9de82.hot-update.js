"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
self["webpackHotUpdate_N_E"]("app/customers/page",{

/***/ "(app-pages-browser)/./app/components/tables/SimpleTable.tsx":
/*!***********************************************!*\
  !*** ./app/components/tables/SimpleTable.tsx ***!
  \***********************************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"(app-pages-browser)/./node_modules/next/dist/compiled/react/jsx-dev-runtime.js\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-OA6OURRG.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-GEJVU65N.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-DRZNIHMG.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-GIQFRSD6.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-MGVPL3OH.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-J4QO5GAJ.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-T2WCTPDH.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_9__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/button/dist/chunk-UVUR7MCU.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_10__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-B5H2YLEF.mjs\");\n/* harmony import */ var _iconify_react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @iconify/react */ \"(app-pages-browser)/./node_modules/@iconify/react/dist/iconify.mjs\");\n/* __next_internal_client_entry_do_not_use__ default auto */ \n\n\nconst SimpleTable = (param)=>{\n    let { columns, data, size = \"md\", showEdit = false, showDelete = false } = param;\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_2__.TableContainer, {\n        children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_3__.Table, {\n            size: size,\n            children: [\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_4__.Thead, {\n                    children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Tr, {\n                        children: [\n                            columns.map((col, index)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {\n                                    children: col.name\n                                }, index, false, {\n                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                    lineNumber: 39,\n                                    columnNumber: 15\n                                }, undefined)),\n                            (showEdit || showDelete) && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {}, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 45,\n                                columnNumber: 15\n                            }, undefined)\n                        ]\n                    }, void 0, true, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                        lineNumber: 37,\n                        columnNumber: 11\n                    }, undefined)\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 36,\n                    columnNumber: 9\n                }, undefined),\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_7__.Tbody, {\n                    children: data.map((dataVal, dataIndex)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Tr, {\n                            children: [\n                                columns.map((col, colIndex)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_8__.Td, {\n                                        children: dataVal[col.key]\n                                    }, colIndex, false, {\n                                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                        lineNumber: 55,\n                                        columnNumber: 17\n                                    }, undefined)),\n                                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_8__.Td, {\n                                    children: [\n                                        showEdit && showDelete && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n                                            className: \"flex\",\n                                            children: [\n                                                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_9__.Button, {\n                                                    size: \"sm\",\n                                                    variant: \"main\",\n                                                    className: \"mr-1\",\n                                                    children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_iconify_react__WEBPACK_IMPORTED_MODULE_1__.Icon, {\n                                                        icon: \"lucide:edit\"\n                                                    }, void 0, false, {\n                                                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                                        lineNumber: 65,\n                                                        columnNumber: 23\n                                                    }, undefined)\n                                                }, void 0, false, {\n                                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                                    lineNumber: 64,\n                                                    columnNumber: 21\n                                                }, undefined),\n                                                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_9__.Button, {\n                                                    size: \"sm\",\n                                                    variant: \"third\",\n                                                    children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_iconify_react__WEBPACK_IMPORTED_MODULE_1__.Icon, {\n                                                        icon: \"wpf:delete\"\n                                                    }, void 0, false, {\n                                                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                                        lineNumber: 68,\n                                                        columnNumber: 23\n                                                    }, undefined)\n                                                }, void 0, false, {\n                                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                                    lineNumber: 67,\n                                                    columnNumber: 21\n                                                }, undefined)\n                                            ]\n                                        }, void 0, true, {\n                                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                            lineNumber: 63,\n                                            columnNumber: 19\n                                        }, undefined),\n                                        showEdit && !showDelete && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_9__.Button, {\n                                            size: \"sm\",\n                                            variant: \"main\",\n                                            children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_iconify_react__WEBPACK_IMPORTED_MODULE_1__.Icon, {\n                                                icon: \"lucide:edit\"\n                                            }, void 0, false, {\n                                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                                lineNumber: 75,\n                                                columnNumber: 21\n                                            }, undefined)\n                                        }, void 0, false, {\n                                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                            lineNumber: 74,\n                                            columnNumber: 19\n                                        }, undefined),\n                                        !showEdit && showDelete && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_9__.Button, {\n                                            size: \"sm\",\n                                            variant: \"third\",\n                                            children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_iconify_react__WEBPACK_IMPORTED_MODULE_1__.Icon, {\n                                                icon: \"wpf:delete\"\n                                            }, void 0, false, {\n                                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                                lineNumber: 81,\n                                                columnNumber: 21\n                                            }, undefined)\n                                        }, void 0, false, {\n                                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                            lineNumber: 80,\n                                            columnNumber: 19\n                                        }, undefined)\n                                    ]\n                                }, void 0, true, {\n                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                    lineNumber: 61,\n                                    columnNumber: 15\n                                }, undefined)\n                            ]\n                        }, dataIndex, true, {\n                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                            lineNumber: 52,\n                            columnNumber: 13\n                        }, undefined))\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 49,\n                    columnNumber: 9\n                }, undefined),\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_10__.Tfoot, {\n                    children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Tr, {\n                        children: [\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {\n                                children: \"To convert\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 93,\n                                columnNumber: 13\n                            }, undefined),\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {\n                                children: \"into\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 94,\n                                columnNumber: 13\n                            }, undefined),\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Th, {\n                                isNumeric: true,\n                                children: \"multiply by\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 95,\n                                columnNumber: 13\n                            }, undefined)\n                        ]\n                    }, void 0, true, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                        lineNumber: 92,\n                        columnNumber: 11\n                    }, undefined)\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 91,\n                    columnNumber: 9\n                }, undefined)\n            ]\n        }, void 0, true, {\n            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n            lineNumber: 35,\n            columnNumber: 7\n        }, undefined)\n    }, void 0, false, {\n        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n        lineNumber: 34,\n        columnNumber: 5\n    }, undefined);\n};\n_c = SimpleTable;\n/* harmony default export */ __webpack_exports__[\"default\"] = (SimpleTable);\nvar _c;\n$RefreshReg$(_c, \"SimpleTable\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevSignature = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevSignature) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports signature on update so we can compare the boundary\n                // signatures. We avoid saving exports themselves since it causes memory leaks (https://github.com/vercel/next.js/pull/53797)\n                module.hot.dispose(function (data) {\n                    data.prevSignature =\n                        self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports);\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevSignature !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevSignature, self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports))) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevSignature !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiKGFwcC1wYWdlcy1icm93c2VyKS8uL2FwcC9jb21wb25lbnRzL3RhYmxlcy9TaW1wbGVUYWJsZS50c3giLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7OztBQWF5QjtBQUNhO0FBVXRDLE1BQU1VLGNBQTBDO1FBQUMsRUFDL0NDLE9BQU8sRUFDUEMsSUFBSSxFQUNKQyxPQUFPLElBQUksRUFDWEMsV0FBVyxLQUFLLEVBQ2hCQyxhQUFhLEtBQUssRUFDbkI7SUFFQyxxQkFDRSw4REFBQ1IsNERBQWNBO2tCQUNiLDRFQUFDUCxtREFBS0E7WUFBQ2EsTUFBTUE7OzhCQUNYLDhEQUFDWixtREFBS0E7OEJBQ0osNEVBQUNHLGdEQUFFQTs7NEJBQ0FPLFFBQVFLLEdBQUcsQ0FBQyxDQUFDQyxLQUF3QkMsc0JBQ3BDLDhEQUFDYixnREFBRUE7OENBQ0FZLElBQUlFLElBQUk7bUNBREZEOzs7Ozs0QkFLVEosQ0FBQUEsWUFBWUMsVUFBUyxtQkFDckIsOERBQUNWLGdEQUFFQTs7Ozs7Ozs7Ozs7Ozs7Ozs4QkFJVCw4REFBQ0gsbURBQUtBOzhCQUNIVSxLQUFLSSxHQUFHLENBQUMsQ0FBQ0ksU0FBY0MsMEJBRXZCLDhEQUFDakIsZ0RBQUVBOztnQ0FDQU8sUUFBUUssR0FBRyxDQUFDLENBQUNDLEtBQXdCSyx5QkFFcEMsOERBQUNoQixnREFBRUE7a0RBQ0FjLE9BQU8sQ0FBQ0gsSUFBSU0sR0FBRyxDQUFDO3VDQURWRDs7Ozs7OENBTVgsOERBQUNoQixnREFBRUE7O3dDQUNBUSxZQUFZQyw0QkFDWCw4REFBQ1M7NENBQUlDLFdBQVU7OzhEQUNiLDhEQUFDakIsb0RBQU1BO29EQUFDSyxNQUFLO29EQUFLYSxTQUFRO29EQUFPRCxXQUFVOzhEQUN6Qyw0RUFBQ2hCLGdEQUFJQTt3REFBQ2tCLE1BQUs7Ozs7Ozs7Ozs7OzhEQUViLDhEQUFDbkIsb0RBQU1BO29EQUFDSyxNQUFLO29EQUFLYSxTQUFROzhEQUN4Qiw0RUFBQ2pCLGdEQUFJQTt3REFBQ2tCLE1BQUs7Ozs7Ozs7Ozs7Ozs7Ozs7O3dDQUtoQmIsWUFBWSxDQUFDQyw0QkFDWiw4REFBQ1Asb0RBQU1BOzRDQUFDSyxNQUFLOzRDQUFLYSxTQUFRO3NEQUN4Qiw0RUFBQ2pCLGdEQUFJQTtnREFBQ2tCLE1BQUs7Ozs7Ozs7Ozs7O3dDQUlkLENBQUNiLFlBQVlDLDRCQUNaLDhEQUFDUCxvREFBTUE7NENBQUNLLE1BQUs7NENBQUthLFNBQVE7c0RBQ3hCLDRFQUFDakIsZ0RBQUlBO2dEQUFDa0IsTUFBSzs7Ozs7Ozs7Ozs7Ozs7Ozs7OzJCQTdCVk47Ozs7Ozs7Ozs7OEJBdUNiLDhEQUFDbEIsb0RBQUtBOzhCQUNKLDRFQUFDQyxnREFBRUE7OzBDQUNELDhEQUFDQyxnREFBRUE7MENBQUM7Ozs7OzswQ0FDSiw4REFBQ0EsZ0RBQUVBOzBDQUFDOzs7Ozs7MENBQ0osOERBQUNBLGdEQUFFQTtnQ0FBQ3VCLFNBQVM7MENBQUM7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7QUFNMUI7S0E1RU1sQjtBQThFTiwrREFBZUEsV0FBV0EsRUFBQyIsInNvdXJjZXMiOlsid2VicGFjazovL19OX0UvLi9hcHAvY29tcG9uZW50cy90YWJsZXMvU2ltcGxlVGFibGUudHN4P2FiNGUiXSwic291cmNlc0NvbnRlbnQiOlsiJ3VzZSBjbGllbnQnO1xuXG5pbXBvcnQgeyBTaW1wbGVUYWJsZUNvbHVtbiB9IGZyb20gJy4vU2ltcGxlVGFibGUudHlwZXMnO1xuaW1wb3J0IHtcbiAgVGFibGUsXG4gIFRoZWFkLFxuICBUYm9keSxcbiAgVGZvb3QsXG4gIFRyLFxuICBUaCxcbiAgVGQsXG4gIFRhYmxlQ29udGFpbmVyLFxuICBCdXR0b24sXG59IGZyb20gJ0BjaGFrcmEtdWkvcmVhY3QnXG5pbXBvcnQgeyBJY29uIH0gZnJvbSAnQGljb25pZnkvcmVhY3QnO1xuXG5pbnRlcmZhY2UgU2ltcGxlVGFibGVQcm9wcyB7XG4gIGNvbHVtbnM6IFNpbXBsZVRhYmxlQ29sdW1uW107XG4gIGRhdGE6IGFueVtdO1xuICBzaXplPzogc3RyaW5nO1xuICBzaG93RWRpdD86IGJvb2xlYW47XG4gIHNob3dEZWxldGU/OiBib29sZWFuO1xufVxuXG5jb25zdCBTaW1wbGVUYWJsZTogUmVhY3QuRkM8U2ltcGxlVGFibGVQcm9wcz4gPSAoe1xuICBjb2x1bW5zLFxuICBkYXRhLFxuICBzaXplID0gJ21kJyxcbiAgc2hvd0VkaXQgPSBmYWxzZSxcbiAgc2hvd0RlbGV0ZSA9IGZhbHNlXG59KSA9PiB7XG5cbiAgcmV0dXJuIChcbiAgICA8VGFibGVDb250YWluZXI+XG4gICAgICA8VGFibGUgc2l6ZT17c2l6ZX0+XG4gICAgICAgIDxUaGVhZD5cbiAgICAgICAgICA8VHI+XG4gICAgICAgICAgICB7Y29sdW1ucy5tYXAoKGNvbDogU2ltcGxlVGFibGVDb2x1bW4sIGluZGV4OiBudW1iZXIpID0+IChcbiAgICAgICAgICAgICAgPFRoIGtleT17aW5kZXh9PlxuICAgICAgICAgICAgICAgIHtjb2wubmFtZX1cbiAgICAgICAgICAgICAgPC9UaD5cbiAgICAgICAgICAgICkpfVxuXG4gICAgICAgICAgICB7KHNob3dFZGl0IHx8IHNob3dEZWxldGUpICYmXG4gICAgICAgICAgICAgIDxUaD48L1RoPlxuICAgICAgICAgICAgfVxuICAgICAgICAgIDwvVHI+XG4gICAgICAgIDwvVGhlYWQ+XG4gICAgICAgIDxUYm9keT5cbiAgICAgICAgICB7ZGF0YS5tYXAoKGRhdGFWYWw6IGFueSwgZGF0YUluZGV4OiBudW1iZXIpID0+IChcblxuICAgICAgICAgICAgPFRyIGtleT17ZGF0YUluZGV4fT5cbiAgICAgICAgICAgICAge2NvbHVtbnMubWFwKChjb2w6IFNpbXBsZVRhYmxlQ29sdW1uLCBjb2xJbmRleDogbnVtYmVyKSA9PiAoXG5cbiAgICAgICAgICAgICAgICA8VGQga2V5PXtjb2xJbmRleH0+XG4gICAgICAgICAgICAgICAgICB7ZGF0YVZhbFtjb2wua2V5XX1cbiAgICAgICAgICAgICAgICA8L1RkPlxuXG4gICAgICAgICAgICAgICkpfVxuXG4gICAgICAgICAgICAgIDxUZD5cbiAgICAgICAgICAgICAgICB7c2hvd0VkaXQgJiYgc2hvd0RlbGV0ZSAmJiAoXG4gICAgICAgICAgICAgICAgICA8ZGl2IGNsYXNzTmFtZT0nZmxleCc+XG4gICAgICAgICAgICAgICAgICAgIDxCdXR0b24gc2l6ZT1cInNtXCIgdmFyaWFudD1cIm1haW5cIiBjbGFzc05hbWU9XCJtci0xXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPEljb24gaWNvbj1cImx1Y2lkZTplZGl0XCIgLz5cbiAgICAgICAgICAgICAgICAgICAgPC9CdXR0b24+XG4gICAgICAgICAgICAgICAgICAgIDxCdXR0b24gc2l6ZT1cInNtXCIgdmFyaWFudD1cInRoaXJkXCI+XG4gICAgICAgICAgICAgICAgICAgICAgPEljb24gaWNvbj1cIndwZjpkZWxldGVcIiAvPlxuICAgICAgICAgICAgICAgICAgICA8L0J1dHRvbj5cbiAgICAgICAgICAgICAgICAgIDwvZGl2PlxuICAgICAgICAgICAgICAgICl9XG5cbiAgICAgICAgICAgICAgICB7c2hvd0VkaXQgJiYgIXNob3dEZWxldGUgJiYgKFxuICAgICAgICAgICAgICAgICAgPEJ1dHRvbiBzaXplPVwic21cIiB2YXJpYW50PVwibWFpblwiPlxuICAgICAgICAgICAgICAgICAgICA8SWNvbiBpY29uPVwibHVjaWRlOmVkaXRcIiAvPlxuICAgICAgICAgICAgICAgICAgPC9CdXR0b24+XG4gICAgICAgICAgICAgICAgKX1cblxuICAgICAgICAgICAgICAgIHshc2hvd0VkaXQgJiYgc2hvd0RlbGV0ZSAmJiAoXG4gICAgICAgICAgICAgICAgICA8QnV0dG9uIHNpemU9XCJzbVwiIHZhcmlhbnQ9XCJ0aGlyZFwiPlxuICAgICAgICAgICAgICAgICAgICA8SWNvbiBpY29uPVwid3BmOmRlbGV0ZVwiIC8+XG4gICAgICAgICAgICAgICAgICA8L0J1dHRvbj5cbiAgICAgICAgICAgICAgICApfVxuICAgICAgICAgICAgICA8L1RkPlxuXG4gICAgICAgICAgICA8L1RyPlxuXG4gICAgICAgICAgKSl9XG5cbiAgICAgICAgPC9UYm9keT5cbiAgICAgICAgPFRmb290PlxuICAgICAgICAgIDxUcj5cbiAgICAgICAgICAgIDxUaD5UbyBjb252ZXJ0PC9UaD5cbiAgICAgICAgICAgIDxUaD5pbnRvPC9UaD5cbiAgICAgICAgICAgIDxUaCBpc051bWVyaWM+bXVsdGlwbHkgYnk8L1RoPlxuICAgICAgICAgIDwvVHI+XG4gICAgICAgIDwvVGZvb3Q+XG4gICAgICA8L1RhYmxlPlxuICAgIDwvVGFibGVDb250YWluZXI+XG4gICk7XG59XG5cbmV4cG9ydCBkZWZhdWx0IFNpbXBsZVRhYmxlO1xuIl0sIm5hbWVzIjpbIlRhYmxlIiwiVGhlYWQiLCJUYm9keSIsIlRmb290IiwiVHIiLCJUaCIsIlRkIiwiVGFibGVDb250YWluZXIiLCJCdXR0b24iLCJJY29uIiwiU2ltcGxlVGFibGUiLCJjb2x1bW5zIiwiZGF0YSIsInNpemUiLCJzaG93RWRpdCIsInNob3dEZWxldGUiLCJtYXAiLCJjb2wiLCJpbmRleCIsIm5hbWUiLCJkYXRhVmFsIiwiZGF0YUluZGV4IiwiY29sSW5kZXgiLCJrZXkiLCJkaXYiLCJjbGFzc05hbWUiLCJ2YXJpYW50IiwiaWNvbiIsImlzTnVtZXJpYyJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///(app-pages-browser)/./app/components/tables/SimpleTable.tsx\n"));

/***/ })

});
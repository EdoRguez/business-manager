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

eval(__webpack_require__.ts("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-dev-runtime */ \"(app-pages-browser)/./node_modules/next/dist/compiled/react/jsx-dev-runtime.js\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-OA6OURRG.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-GEJVU65N.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-DRZNIHMG.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-GIQFRSD6.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-MGVPL3OH.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-J4QO5GAJ.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_7__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-T2WCTPDH.mjs\");\n/* harmony import */ var _chakra_ui_react__WEBPACK_IMPORTED_MODULE_8__ = __webpack_require__(/*! @chakra-ui/react */ \"(app-pages-browser)/./node_modules/@chakra-ui/table/dist/chunk-B5H2YLEF.mjs\");\n/* __next_internal_client_entry_do_not_use__ default auto */ \n\nconst SimpleTable = (param)=>{\n    let { columns, data, size = \"md\", showEdit = false, showDelete = false } = param;\n    const btnEditDelete = ()=>{\n        return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(\"div\", {\n            children: \"Hola\"\n        }, void 0, false, {\n            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n            lineNumber: 33,\n            columnNumber: 7\n        }, undefined);\n    };\n    return /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_1__.TableContainer, {\n        children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_2__.Table, {\n            size: size,\n            children: [\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_3__.Thead, {\n                    children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_4__.Tr, {\n                        children: [\n                            columns.map((col, index)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Th, {\n                                    children: col.name\n                                }, index, false, {\n                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                    lineNumber: 45,\n                                    columnNumber: 15\n                                }, undefined)),\n                            (showEdit || showDelete) && /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Th, {}, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 51,\n                                columnNumber: 15\n                            }, undefined)\n                        ]\n                    }, void 0, true, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                        lineNumber: 43,\n                        columnNumber: 11\n                    }, undefined)\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 42,\n                    columnNumber: 9\n                }, undefined),\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_6__.Tbody, {\n                    children: data.map((dataVal, dataIndex)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_4__.Tr, {\n                            children: [\n                                columns.map((col, colIndex)=>/*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_7__.Td, {\n                                        children: dataVal[col.key]\n                                    }, colIndex, false, {\n                                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                        lineNumber: 61,\n                                        columnNumber: 17\n                                    }, undefined)),\n                                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_7__.Td, {\n                                    children: btnEditDelete\n                                }, 9999, false, {\n                                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                    lineNumber: 67,\n                                    columnNumber: 15\n                                }, undefined)\n                            ]\n                        }, dataIndex, true, {\n                            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                            lineNumber: 58,\n                            columnNumber: 13\n                        }, undefined))\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 55,\n                    columnNumber: 9\n                }, undefined),\n                /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_8__.Tfoot, {\n                    children: /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_4__.Tr, {\n                        children: [\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Th, {\n                                children: \"To convert\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 78,\n                                columnNumber: 13\n                            }, undefined),\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Th, {\n                                children: \"into\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 79,\n                                columnNumber: 13\n                            }, undefined),\n                            /*#__PURE__*/ (0,react_jsx_dev_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxDEV)(_chakra_ui_react__WEBPACK_IMPORTED_MODULE_5__.Th, {\n                                isNumeric: true,\n                                children: \"multiply by\"\n                            }, void 0, false, {\n                                fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                                lineNumber: 80,\n                                columnNumber: 13\n                            }, undefined)\n                        ]\n                    }, void 0, true, {\n                        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                        lineNumber: 77,\n                        columnNumber: 11\n                    }, undefined)\n                }, void 0, false, {\n                    fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n                    lineNumber: 76,\n                    columnNumber: 9\n                }, undefined)\n            ]\n        }, void 0, true, {\n            fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n            lineNumber: 41,\n            columnNumber: 7\n        }, undefined)\n    }, void 0, false, {\n        fileName: \"C:\\\\Users\\\\User\\\\Documents\\\\Projects\\\\business-manager\\\\src\\\\ui\\\\app\\\\components\\\\tables\\\\SimpleTable.tsx\",\n        lineNumber: 40,\n        columnNumber: 5\n    }, undefined);\n};\n_c = SimpleTable;\n/* harmony default export */ __webpack_exports__[\"default\"] = (SimpleTable);\nvar _c;\n$RefreshReg$(_c, \"SimpleTable\");\n\n\n;\n    // Wrapped in an IIFE to avoid polluting the global scope\n    ;\n    (function () {\n        var _a, _b;\n        // Legacy CSS implementations will `eval` browser code in a Node.js context\n        // to extract CSS. For backwards compatibility, we need to check we're in a\n        // browser context before continuing.\n        if (typeof self !== 'undefined' &&\n            // AMP / No-JS mode does not inject these helpers:\n            '$RefreshHelpers$' in self) {\n            // @ts-ignore __webpack_module__ is global\n            var currentExports = module.exports;\n            // @ts-ignore __webpack_module__ is global\n            var prevSignature = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevSignature) !== null && _b !== void 0 ? _b : null;\n            // This cannot happen in MainTemplate because the exports mismatch between\n            // templating and execution.\n            self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n            // A module can be accepted automatically based on its exports, e.g. when\n            // it is a Refresh Boundary.\n            if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n                // Save the previous exports signature on update so we can compare the boundary\n                // signatures. We avoid saving exports themselves since it causes memory leaks (https://github.com/vercel/next.js/pull/53797)\n                module.hot.dispose(function (data) {\n                    data.prevSignature =\n                        self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports);\n                });\n                // Unconditionally accept an update to this module, we'll check if it's\n                // still a Refresh Boundary later.\n                // @ts-ignore importMeta is replaced in the loader\n                module.hot.accept();\n                // This field is set when the previous version of this module was a\n                // Refresh Boundary, letting us know we need to check for invalidation or\n                // enqueue an update.\n                if (prevSignature !== null) {\n                    // A boundary can become ineligible if its exports are incompatible\n                    // with the previous exports.\n                    //\n                    // For example, if you add/remove/change exports, we'll want to\n                    // re-execute the importing modules, and force those components to\n                    // re-render. Similarly, if you convert a class component to a\n                    // function, we want to invalidate the boundary.\n                    if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevSignature, self.$RefreshHelpers$.getRefreshBoundarySignature(currentExports))) {\n                        module.hot.invalidate();\n                    }\n                    else {\n                        self.$RefreshHelpers$.scheduleUpdate();\n                    }\n                }\n            }\n            else {\n                // Since we just executed the code for the module, it's possible that the\n                // new exports made it ineligible for being a boundary.\n                // We only care about the case when we were _previously_ a boundary,\n                // because we already accepted this update (accidental side effect).\n                var isNoLongerABoundary = prevSignature !== null;\n                if (isNoLongerABoundary) {\n                    module.hot.invalidate();\n                }\n            }\n        }\n    })();\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiKGFwcC1wYWdlcy1icm93c2VyKS8uL2FwcC9jb21wb25lbnRzL3RhYmxlcy9TaW1wbGVUYWJsZS50c3giLCJtYXBwaW5ncyI6Ijs7Ozs7Ozs7Ozs7QUFZeUI7QUFVekIsTUFBTVEsY0FBMEM7UUFBQyxFQUMvQ0MsT0FBTyxFQUNQQyxJQUFJLEVBQ0pDLE9BQU8sSUFBSSxFQUNYQyxXQUFXLEtBQUssRUFDaEJDLGFBQWEsS0FBSyxFQUNuQjtJQUVDLE1BQU1DLGdCQUF5QjtRQUM3QixxQkFDRSw4REFBQ0M7c0JBQUk7Ozs7OztJQUlUO0lBRUEscUJBQ0UsOERBQUNSLDREQUFjQTtrQkFDYiw0RUFBQ1AsbURBQUtBO1lBQUNXLE1BQU1BOzs4QkFDWCw4REFBQ1YsbURBQUtBOzhCQUNKLDRFQUFDRyxnREFBRUE7OzRCQUNBSyxRQUFRTyxHQUFHLENBQUMsQ0FBQ0MsS0FBd0JDLHNCQUNwQyw4REFBQ2IsZ0RBQUVBOzhDQUNBWSxJQUFJRSxJQUFJO21DQURGRDs7Ozs7NEJBS1ROLENBQUFBLFlBQVlDLFVBQVMsbUJBQ3JCLDhEQUFDUixnREFBRUE7Ozs7Ozs7Ozs7Ozs7Ozs7OEJBSVQsOERBQUNILG1EQUFLQTs4QkFDSFEsS0FBS00sR0FBRyxDQUFDLENBQUNJLFNBQWNDLDBCQUV2Qiw4REFBQ2pCLGdEQUFFQTs7Z0NBQ0FLLFFBQVFPLEdBQUcsQ0FBQyxDQUFDQyxLQUF3QksseUJBRXBDLDhEQUFDaEIsZ0RBQUVBO2tEQUNBYyxPQUFPLENBQUNILElBQUlNLEdBQUcsQ0FBQzt1Q0FEVkQ7Ozs7OzhDQU1YLDhEQUFDaEIsZ0RBQUVBOzhDQUNBUTttQ0FETTs7Ozs7OzJCQVRGTzs7Ozs7Ozs7Ozs4QkFrQmIsOERBQUNsQixtREFBS0E7OEJBQ0osNEVBQUNDLGdEQUFFQTs7MENBQ0QsOERBQUNDLGdEQUFFQTswQ0FBQzs7Ozs7OzBDQUNKLDhEQUFDQSxnREFBRUE7MENBQUM7Ozs7OzswQ0FDSiw4REFBQ0EsZ0RBQUVBO2dDQUFDbUIsU0FBUzswQ0FBQzs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7Ozs7OztBQU0xQjtLQS9ETWhCO0FBaUVOLCtEQUFlQSxXQUFXQSxFQUFDIiwic291cmNlcyI6WyJ3ZWJwYWNrOi8vX05fRS8uL2FwcC9jb21wb25lbnRzL3RhYmxlcy9TaW1wbGVUYWJsZS50c3g/YWI0ZSJdLCJzb3VyY2VzQ29udGVudCI6WyIndXNlIGNsaWVudCc7XG5cbmltcG9ydCB7IFNpbXBsZVRhYmxlQ29sdW1uIH0gZnJvbSAnLi9TaW1wbGVUYWJsZS50eXBlcyc7XG5pbXBvcnQge1xuICBUYWJsZSxcbiAgVGhlYWQsXG4gIFRib2R5LFxuICBUZm9vdCxcbiAgVHIsXG4gIFRoLFxuICBUZCxcbiAgVGFibGVDb250YWluZXIsXG59IGZyb20gJ0BjaGFrcmEtdWkvcmVhY3QnXG5cbmludGVyZmFjZSBTaW1wbGVUYWJsZVByb3BzIHtcbiAgY29sdW1uczogU2ltcGxlVGFibGVDb2x1bW5bXTtcbiAgZGF0YTogYW55W107XG4gIHNpemU/OiBzdHJpbmc7XG4gIHNob3dFZGl0PzogYm9vbGVhbjtcbiAgc2hvd0RlbGV0ZT86IGJvb2xlYW47XG59XG5cbmNvbnN0IFNpbXBsZVRhYmxlOiBSZWFjdC5GQzxTaW1wbGVUYWJsZVByb3BzPiA9ICh7XG4gIGNvbHVtbnMsXG4gIGRhdGEsXG4gIHNpemUgPSAnbWQnLFxuICBzaG93RWRpdCA9IGZhbHNlLFxuICBzaG93RGVsZXRlID0gZmFsc2Vcbn0pID0+IHtcbiAgXG4gIGNvbnN0IGJ0bkVkaXREZWxldGU6IEVsZW1lbnQgPSAoKSA9PiB7XG4gICAgcmV0dXJuIChcbiAgICAgIDxkaXY+XG4gICAgICAgIEhvbGFcbiAgICAgIDwvZGl2PlxuICAgIClcbiAgfTtcblxuICByZXR1cm4gKFxuICAgIDxUYWJsZUNvbnRhaW5lcj5cbiAgICAgIDxUYWJsZSBzaXplPXtzaXplfT5cbiAgICAgICAgPFRoZWFkPlxuICAgICAgICAgIDxUcj5cbiAgICAgICAgICAgIHtjb2x1bW5zLm1hcCgoY29sOiBTaW1wbGVUYWJsZUNvbHVtbiwgaW5kZXg6IG51bWJlcikgPT4gKFxuICAgICAgICAgICAgICA8VGgga2V5PXtpbmRleH0+XG4gICAgICAgICAgICAgICAge2NvbC5uYW1lfVxuICAgICAgICAgICAgICA8L1RoPlxuICAgICAgICAgICAgKSl9XG5cbiAgICAgICAgICAgIHsoc2hvd0VkaXQgfHwgc2hvd0RlbGV0ZSkgJiZcbiAgICAgICAgICAgICAgPFRoPjwvVGg+XG4gICAgICAgICAgICB9XG4gICAgICAgICAgPC9Ucj5cbiAgICAgICAgPC9UaGVhZD5cbiAgICAgICAgPFRib2R5PlxuICAgICAgICAgIHtkYXRhLm1hcCgoZGF0YVZhbDogYW55LCBkYXRhSW5kZXg6IG51bWJlcikgPT4gKFxuXG4gICAgICAgICAgICA8VHIga2V5PXtkYXRhSW5kZXh9PlxuICAgICAgICAgICAgICB7Y29sdW1ucy5tYXAoKGNvbDogU2ltcGxlVGFibGVDb2x1bW4sIGNvbEluZGV4OiBudW1iZXIpID0+IChcblxuICAgICAgICAgICAgICAgIDxUZCBrZXk9e2NvbEluZGV4fT5cbiAgICAgICAgICAgICAgICAgIHtkYXRhVmFsW2NvbC5rZXldfVxuICAgICAgICAgICAgICAgIDwvVGQ+XG5cbiAgICAgICAgICAgICAgKSl9XG4gICAgICAgICAgICAgIFxuICAgICAgICAgICAgICA8VGQga2V5PXs5OTk5fT5cbiAgICAgICAgICAgICAgICB7YnRuRWRpdERlbGV0ZX1cbiAgICAgICAgICAgICAgPC9UZD5cblxuICAgICAgICAgICAgPC9Ucj5cblxuICAgICAgICAgICkpfVxuXG4gICAgICAgIDwvVGJvZHk+XG4gICAgICAgIDxUZm9vdD5cbiAgICAgICAgICA8VHI+XG4gICAgICAgICAgICA8VGg+VG8gY29udmVydDwvVGg+XG4gICAgICAgICAgICA8VGg+aW50bzwvVGg+XG4gICAgICAgICAgICA8VGggaXNOdW1lcmljPm11bHRpcGx5IGJ5PC9UaD5cbiAgICAgICAgICA8L1RyPlxuICAgICAgICA8L1Rmb290PlxuICAgICAgPC9UYWJsZT5cbiAgICA8L1RhYmxlQ29udGFpbmVyPlxuICApO1xufVxuXG5leHBvcnQgZGVmYXVsdCBTaW1wbGVUYWJsZTtcbiJdLCJuYW1lcyI6WyJUYWJsZSIsIlRoZWFkIiwiVGJvZHkiLCJUZm9vdCIsIlRyIiwiVGgiLCJUZCIsIlRhYmxlQ29udGFpbmVyIiwiU2ltcGxlVGFibGUiLCJjb2x1bW5zIiwiZGF0YSIsInNpemUiLCJzaG93RWRpdCIsInNob3dEZWxldGUiLCJidG5FZGl0RGVsZXRlIiwiZGl2IiwibWFwIiwiY29sIiwiaW5kZXgiLCJuYW1lIiwiZGF0YVZhbCIsImRhdGFJbmRleCIsImNvbEluZGV4Iiwia2V5IiwiaXNOdW1lcmljIl0sInNvdXJjZVJvb3QiOiIifQ==\n//# sourceURL=webpack-internal:///(app-pages-browser)/./app/components/tables/SimpleTable.tsx\n"));

/***/ })

});
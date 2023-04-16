<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

Route::post('/login', [App\Http\Controllers\UserController::class, 'login']);

Route::middleware(['auth:sanctum', 'admin'])->group(function() {
    Route::get('/admin/loans', [App\Http\Controllers\AdminController::class, 'loans']);
    Route::post('/admin/loans/{loan}/approve', [App\Http\Controllers\AdminController::class, 'approve'])->missing(function (Request $request) {
        return response()->json([
            'status' => 'failed',
            'message' => 'Loan not found'
        ], 404);
    });
});

Route::middleware('auth:sanctum')->resource('loans', App\Http\Controllers\LoanController::class)->missing(function (Request $request) {
    return response()->json([
        'status' => 'failed',
        'message' => 'Loan not found'
    ], 404);
});

Route::middleware('auth:sanctum')->get('/loans/{loan}/repayments', [App\Http\Controllers\LoanController::class, 'show_repayments'])->missing(function (Request $request) {
    return response()->json([
        'status' => 'failed',
        'message' => 'Loan not found'
    ], 404);
});

Route::middleware('auth:sanctum')->post('/loans/{loan}/repayments', [App\Http\Controllers\LoanController::class, 'pay_repayments'])->missing(function (Request $request) {
    return response()->json([
        'status' => 'failed',
        'message' => 'Loan not found'
    ], 404);
});

Route::get('/health', function() {
    return 'ok';
});

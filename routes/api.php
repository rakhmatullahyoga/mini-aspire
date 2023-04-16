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
});

Route::middleware('auth:sanctum')->resource('loans', App\Http\Controllers\LoanController::class);

<?php

use Illuminate\Support\Facades\Route;
use App\Http\Controllers\MerchantController;

Route::get('/get-merchant', [MerchantController::class, 'getMerchant']);

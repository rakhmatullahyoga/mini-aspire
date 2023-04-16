<?php

namespace App\Http\Controllers;

use App\Models\Loan;
use Illuminate\Http\Request;

class AdminController extends Controller
{
    public function loans(Request $request)
    {
        $loans = Loan::latest()->paginate(10);
        return response()->json([
            'status' => 'success',
            'data' => $loans
        ]);
    }
}

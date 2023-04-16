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

    public function approve(Loan $loan)
    {
        $loan->status = 'approved';
        $loan->save();
        return response()->json([
            'status' => 'success',
            'data' => $loan
        ]);
    }
}

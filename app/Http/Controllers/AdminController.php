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

    public function approve(string $id)
    {
        $loan = Loan::find($id);
        if ($loan == null) {
            return response()->json([
                'status' => 'failed',
                'message' => 'Loan not found'
            ]);
        }
        $loan->status = 'approved';
        $loan->save();
        return response()->json([
            'status' => 'success',
            'data' => $loan
        ]);
    }
}

<?php

namespace App\Http\Controllers;

use App\Models\Loan;
use App\Models\Repayment;
use DateTime;
use Illuminate\Http\Request;

class LoanController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index(Request $request)
    {
        $loans = $request->user()->loans()->latest()->paginate(10);
        return response()->json([
            'status' => 'success',
            'data' => $loans
        ]);
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        $input = $request->all();
        $input['user_id'] = $request->user()->id;
        $input['status'] = 'pending';
        $loan = Loan::create($input);
        $repayment_amount = $input['amount'] / (float) $input['term'];
        $due_date = new DateTime($input['loan_date']);
        for ($i=0; $i<$input['term']; $i++) {
            $due_date->modify('+7 day');
            $repayment = new Repayment;
            $repayment->loan_id = $loan->id;
            $repayment->status = 'pending';
            $repayment->amount = $repayment_amount;
            $repayment->due_date = $due_date;
            $repayment->save();
        }
        return response()->json([
            'status' => 'success',
            'data' => $loan
        ]);
    }

    /**
     * Display the specified resource.
     */
    public function show(Request $request, Loan $loan)
    {
        if ($request->user()->cannot('view', $loan)) {
            return response()->json([
                'status' => 'failed',
                'message' => 'Cannot find your loan'
            ], 404);
        }
        return response()->json([
            'status' => 'success',
            'data' => $loan
        ]);
    }
}

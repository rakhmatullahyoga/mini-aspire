<?php

namespace Database\Seeders;

// use Illuminate\Database\Console\Seeds\WithoutModelEvents;
use Illuminate\Database\Seeder;

class DatabaseSeeder extends Seeder
{
    /**
     * Seed the application's database.
     */
    public function run(): void
    {
        \App\Models\User::factory()->create([
            'name' => 'Admin',
            'email' => 'admin@mini-aspire.com',
            'password' => bcrypt('admin'),
        ]);

        \App\Models\User::factory()->create([
            'name' => 'Customer',
            'email' => 'customer@mini-aspire.com',
            'password' => bcrypt('customer'),
        ]);
    }
}

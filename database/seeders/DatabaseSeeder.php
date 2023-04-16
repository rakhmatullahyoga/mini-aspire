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
            'name' => 'Customer 1',
            'email' => 'customer1@mini-aspire.com',
            'password' => bcrypt('customer1'),
        ]);

        \App\Models\User::factory()->create([
            'name' => 'Customer 2',
            'email' => 'customer2@mini-aspire.com',
            'password' => bcrypt('customer2'),
        ]);
    }
}

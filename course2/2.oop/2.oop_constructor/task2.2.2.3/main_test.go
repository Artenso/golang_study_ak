package main

import (
	"bytes"
	"os"
	"reflect"
	"testing"
)

func TestNewCustomer(t *testing.T) {
	type args struct {
		id   int
		opts []CustomerOption
	}

	savings := &SavingsAccount{
		balance: 1000,
	}

	tests := []struct {
		name string
		args args
		want *Customer
	}{
		{
			name: "OK",
			args: args{
				id: 1,
				opts: []CustomerOption{
					WithName("John Doe"),
					WithAccount(savings),
				},
			},
			want: &Customer{
				ID:      1,
				Name:    "John Doe",
				Account: savings,
			},
		},
		{
			name: "Without account",
			args: args{
				id: 1,
				opts: []CustomerOption{
					WithName("John Doe"),
				},
			},
			want: &Customer{
				ID:   1,
				Name: "John Doe",
			},
		},
		{
			name: "Without name",
			args: args{
				id: 1,
				opts: []CustomerOption{
					WithAccount(savings),
				},
			},
			want: &Customer{
				ID:      1,
				Account: savings,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCustomer(tt.args.id, tt.args.opts...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCustomer() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavingsAccount_Balance(t *testing.T) {
	tests := []struct {
		name string
		sa   *SavingsAccount
		want float64
	}{
		{
			name: "#1",
			sa: &SavingsAccount{
				balance: 1000,
			},
			want: float64(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.sa.Balance(); got != tt.want {
				t.Errorf("SavingsAccount.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSavingsAccount_Deposit(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name        string
		sa          *SavingsAccount
		args        args
		wantErr     bool
		wantBalance float64
	}{
		{
			name: "OK",
			sa: &SavingsAccount{
				balance: 1000,
			},
			args: args{
				amount: 500,
			},
			wantErr:     false,
			wantBalance: 1500,
		},
		{
			name: "Invalid amount",
			sa: &SavingsAccount{
				balance: 1000,
			},
			args: args{
				amount: -500,
			},
			wantErr:     true,
			wantBalance: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sa.Deposit(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("SavingsAccount.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.sa.balance != tt.wantBalance {
				t.Errorf("Balance after SavingsAccount.Deposit() = %v, want = %v", tt.sa.balance, tt.wantBalance)
			}
		})
	}
}

func TestSavingsAccount_Withdraw(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name        string
		sa          *SavingsAccount
		args        args
		wantErr     bool
		wantBalance float64
	}{
		{
			name: "OK",
			sa: &SavingsAccount{
				balance: 1000,
			},
			args: args{
				amount: 200,
			},
			wantErr:     false,
			wantBalance: 800,
		},
		{
			name: "Invalid amount",
			sa: &SavingsAccount{
				balance: 1000,
			},
			args: args{
				amount: -200,
			},
			wantErr:     true,
			wantBalance: 1000,
		},
		{
			name: "Low balance",
			sa: &SavingsAccount{
				balance: 800,
			},
			args: args{
				amount: 200,
			},
			wantErr:     true,
			wantBalance: 800,
		},
		{
			name: "Not enough money",
			sa: &SavingsAccount{
				balance: 1500,
			},
			args: args{
				amount: 2000,
			},
			wantErr:     true,
			wantBalance: 1500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sa.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("SavingsAccount.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckingAccount_Balance(t *testing.T) {
	tests := []struct {
		name string
		ca   *CheckingAccount
		want float64
	}{
		{
			name: "#1",
			ca: &CheckingAccount{
				balance: 1000,
			},
			want: float64(1000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ca.Balance(); got != tt.want {
				t.Errorf("CheckingAccount.Balance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCheckingAccount_Deposit(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name        string
		ca          *CheckingAccount
		args        args
		wantErr     bool
		wantBalance float64
	}{
		{
			name: "OK",
			ca: &CheckingAccount{
				balance: 1000,
			},
			args: args{
				amount: 500,
			},
			wantErr:     false,
			wantBalance: 1500,
		},
		{
			name: "Invalid amount",
			ca: &CheckingAccount{
				balance: 1000,
			},
			args: args{
				amount: -500,
			},
			wantErr:     true,
			wantBalance: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ca.Deposit(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CheckingAccount.Deposit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheckingAccount_Withdraw(t *testing.T) {
	type args struct {
		amount float64
	}
	tests := []struct {
		name        string
		ca          *CheckingAccount
		args        args
		wantErr     bool
		wantBalance float64
	}{
		{
			name: "OK",
			ca: &CheckingAccount{
				balance: 1000,
			},
			args: args{
				amount: 200,
			},
			wantErr:     false,
			wantBalance: 800,
		},
		{
			name: "Invalid amount",
			ca: &CheckingAccount{
				balance: 1000,
			},
			args: args{
				amount: -200,
			},
			wantErr:     true,
			wantBalance: 1000,
		},
		{
			name: "Not enough money",
			ca: &CheckingAccount{
				balance: 1500,
			},
			args: args{
				amount: 2000,
			},
			wantErr:     true,
			wantBalance: 1500,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ca.Withdraw(tt.args.amount); (err != nil) != tt.wantErr {
				t.Errorf("CheckingAccount.Withdraw() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_main(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		oldOut := os.Stdout

		r, w, err := os.Pipe()
		if err != nil {
			t.Errorf("failed create pipe: %s", err)
		}

		os.Stdout = w

		main()
		w.Close()

		os.Stdout = oldOut

		var out bytes.Buffer
		out.ReadFrom(r)
		expected := "Customer: John Doe, Balance: 900\n"
		if out.String() != expected {
			t.Errorf("want: %s\ngot: %s", expected, out.String())
		}
	})
}

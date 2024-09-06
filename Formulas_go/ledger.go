package main

// Ledger - Command-line, double-entry accounting tool
// Homepage: https://ledger-cli.org/

import (
	"fmt"
	
	"os/exec"
)

func installLedger() {
	// Método 1: Descargar y extraer .tar.gz
	ledger_tar_url := "https://github.com/ledger/ledger/archive/refs/tags/v3.3.2.tar.gz"
	ledger_cmd_tar := exec.Command("curl", "-L", ledger_tar_url, "-o", "package.tar.gz")
	err := ledger_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ledger_zip_url := "https://github.com/ledger/ledger/archive/refs/tags/v3.3.2.zip"
	ledger_cmd_zip := exec.Command("curl", "-L", ledger_zip_url, "-o", "package.zip")
	err = ledger_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ledger_bin_url := "https://github.com/ledger/ledger/archive/refs/tags/v3.3.2.bin"
	ledger_cmd_bin := exec.Command("curl", "-L", ledger_bin_url, "-o", "binary.bin")
	err = ledger_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ledger_src_url := "https://github.com/ledger/ledger/archive/refs/tags/v3.3.2.src.tar.gz"
	ledger_cmd_src := exec.Command("curl", "-L", ledger_src_url, "-o", "source.tar.gz")
	err = ledger_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ledger_cmd_direct := exec.Command("./binary")
	err = ledger_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: texinfo")
exec.Command("latte", "install", "texinfo")
	fmt.Println("Instalando dependencia: boost@1.85")
exec.Command("latte", "install", "boost@1.85")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: gpgme")
exec.Command("latte", "install", "gpgme")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: libassuan")
exec.Command("latte", "install", "libassuan")
}

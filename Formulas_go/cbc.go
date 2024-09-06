package main

// Cbc - Mixed integer linear programming solver
// Homepage: https://github.com/coin-or/Cbc

import (
	"fmt"
	
	"os/exec"
)

func installCbc() {
	// Método 1: Descargar y extraer .tar.gz
	cbc_tar_url := "https://github.com/coin-or/Cbc/archive/refs/tags/releases/2.10.12.tar.gz"
	cbc_cmd_tar := exec.Command("curl", "-L", cbc_tar_url, "-o", "package.tar.gz")
	err := cbc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cbc_zip_url := "https://github.com/coin-or/Cbc/archive/refs/tags/releases/2.10.12.zip"
	cbc_cmd_zip := exec.Command("curl", "-L", cbc_zip_url, "-o", "package.zip")
	err = cbc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cbc_bin_url := "https://github.com/coin-or/Cbc/archive/refs/tags/releases/2.10.12.bin"
	cbc_cmd_bin := exec.Command("curl", "-L", cbc_bin_url, "-o", "binary.bin")
	err = cbc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cbc_src_url := "https://github.com/coin-or/Cbc/archive/refs/tags/releases/2.10.12.src.tar.gz"
	cbc_cmd_src := exec.Command("curl", "-L", cbc_src_url, "-o", "source.tar.gz")
	err = cbc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cbc_cmd_direct := exec.Command("./binary")
	err = cbc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cgl")
exec.Command("latte", "install", "cgl")
	fmt.Println("Instalando dependencia: clp")
exec.Command("latte", "install", "clp")
	fmt.Println("Instalando dependencia: coinutils")
exec.Command("latte", "install", "coinutils")
	fmt.Println("Instalando dependencia: osi")
exec.Command("latte", "install", "osi")
	fmt.Println("Instalando dependencia: openblas")
exec.Command("latte", "install", "openblas")
}

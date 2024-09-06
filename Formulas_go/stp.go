package main

// Stp - Simple Theorem Prover, an efficient SMT solver for bitvectors
// Homepage: https://stp.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installStp() {
	// Método 1: Descargar y extraer .tar.gz
	stp_tar_url := "https://github.com/stp/stp/archive/refs/tags/2.3.4.tar.gz"
	stp_cmd_tar := exec.Command("curl", "-L", stp_tar_url, "-o", "package.tar.gz")
	err := stp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	stp_zip_url := "https://github.com/stp/stp/archive/refs/tags/2.3.4.zip"
	stp_cmd_zip := exec.Command("curl", "-L", stp_zip_url, "-o", "package.zip")
	err = stp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	stp_bin_url := "https://github.com/stp/stp/archive/refs/tags/2.3.4.bin"
	stp_cmd_bin := exec.Command("curl", "-L", stp_bin_url, "-o", "binary.bin")
	err = stp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	stp_src_url := "https://github.com/stp/stp/archive/refs/tags/2.3.4.src.tar.gz"
	stp_cmd_src := exec.Command("curl", "-L", stp_src_url, "-o", "source.tar.gz")
	err = stp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	stp_cmd_direct := exec.Command("./binary")
	err = stp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: bison")
exec.Command("latte", "install", "bison")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: flex")
exec.Command("latte", "install", "flex")
	fmt.Println("Instalando dependencia: python-setuptools")
exec.Command("latte", "install", "python-setuptools")
	fmt.Println("Instalando dependencia: boost")
exec.Command("latte", "install", "boost")
	fmt.Println("Instalando dependencia: cryptominisat")
exec.Command("latte", "install", "cryptominisat")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: minisat")
exec.Command("latte", "install", "minisat")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}

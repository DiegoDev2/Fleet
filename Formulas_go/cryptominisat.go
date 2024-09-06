package main

// Cryptominisat - Advanced SAT solver
// Homepage: https://www.msoos.org/cryptominisat5/

import (
	"fmt"
	
	"os/exec"
)

func installCryptominisat() {
	// Método 1: Descargar y extraer .tar.gz
	cryptominisat_tar_url := "https://github.com/msoos/cryptominisat/archive/refs/tags/5.11.21.tar.gz"
	cryptominisat_cmd_tar := exec.Command("curl", "-L", cryptominisat_tar_url, "-o", "package.tar.gz")
	err := cryptominisat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cryptominisat_zip_url := "https://github.com/msoos/cryptominisat/archive/refs/tags/5.11.21.zip"
	cryptominisat_cmd_zip := exec.Command("curl", "-L", cryptominisat_zip_url, "-o", "package.zip")
	err = cryptominisat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cryptominisat_bin_url := "https://github.com/msoos/cryptominisat/archive/refs/tags/5.11.21.bin"
	cryptominisat_cmd_bin := exec.Command("curl", "-L", cryptominisat_bin_url, "-o", "binary.bin")
	err = cryptominisat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cryptominisat_src_url := "https://github.com/msoos/cryptominisat/archive/refs/tags/5.11.21.src.tar.gz"
	cryptominisat_cmd_src := exec.Command("curl", "-L", cryptominisat_src_url, "-o", "source.tar.gz")
	err = cryptominisat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cryptominisat_cmd_direct := exec.Command("./binary")
	err = cryptominisat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}

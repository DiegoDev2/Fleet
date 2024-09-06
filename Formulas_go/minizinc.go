package main

// Minizinc - Medium-level constraint modeling language
// Homepage: https://www.minizinc.org/

import (
	"fmt"
	
	"os/exec"
)

func installMinizinc() {
	// Método 1: Descargar y extraer .tar.gz
	minizinc_tar_url := "https://github.com/MiniZinc/libminizinc/archive/refs/tags/2.8.5.tar.gz"
	minizinc_cmd_tar := exec.Command("curl", "-L", minizinc_tar_url, "-o", "package.tar.gz")
	err := minizinc_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minizinc_zip_url := "https://github.com/MiniZinc/libminizinc/archive/refs/tags/2.8.5.zip"
	minizinc_cmd_zip := exec.Command("curl", "-L", minizinc_zip_url, "-o", "package.zip")
	err = minizinc_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minizinc_bin_url := "https://github.com/MiniZinc/libminizinc/archive/refs/tags/2.8.5.bin"
	minizinc_cmd_bin := exec.Command("curl", "-L", minizinc_bin_url, "-o", "binary.bin")
	err = minizinc_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minizinc_src_url := "https://github.com/MiniZinc/libminizinc/archive/refs/tags/2.8.5.src.tar.gz"
	minizinc_cmd_src := exec.Command("curl", "-L", minizinc_src_url, "-o", "source.tar.gz")
	err = minizinc_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minizinc_cmd_direct := exec.Command("./binary")
	err = minizinc_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: cbc")
exec.Command("latte", "install", "cbc")
	fmt.Println("Instalando dependencia: cgl")
exec.Command("latte", "install", "cgl")
	fmt.Println("Instalando dependencia: clp")
exec.Command("latte", "install", "clp")
	fmt.Println("Instalando dependencia: coinutils")
exec.Command("latte", "install", "coinutils")
	fmt.Println("Instalando dependencia: gecode")
exec.Command("latte", "install", "gecode")
	fmt.Println("Instalando dependencia: osi")
exec.Command("latte", "install", "osi")
}

package main

// FregeRepl - REPL (read-eval-print loop) for Frege
// Homepage: https://github.com/Frege/frege-repl

import (
	"fmt"
	
	"os/exec"
)

func installFregeRepl() {
	// Método 1: Descargar y extraer .tar.gz
	fregerepl_tar_url := "https://github.com/Frege/frege-repl/releases/download/1.4-SNAPSHOT/frege-repl-1.4-SNAPSHOT.zip"
	fregerepl_cmd_tar := exec.Command("curl", "-L", fregerepl_tar_url, "-o", "package.tar.gz")
	err := fregerepl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fregerepl_zip_url := "https://github.com/Frege/frege-repl/releases/download/1.4-SNAPSHOT/frege-repl-1.4-SNAPSHOT.zip"
	fregerepl_cmd_zip := exec.Command("curl", "-L", fregerepl_zip_url, "-o", "package.zip")
	err = fregerepl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fregerepl_bin_url := "https://github.com/Frege/frege-repl/releases/download/1.4-SNAPSHOT/frege-repl-1.4-SNAPSHOT.zip"
	fregerepl_cmd_bin := exec.Command("curl", "-L", fregerepl_bin_url, "-o", "binary.bin")
	err = fregerepl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fregerepl_src_url := "https://github.com/Frege/frege-repl/releases/download/1.4-SNAPSHOT/frege-repl-1.4-SNAPSHOT.zip"
	fregerepl_cmd_src := exec.Command("curl", "-L", fregerepl_src_url, "-o", "source.tar.gz")
	err = fregerepl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fregerepl_cmd_direct := exec.Command("./binary")
	err = fregerepl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk@17")
exec.Command("latte", "install", "openjdk@17")
}

package main

// Leiningen - Build tool for Clojure
// Homepage: https://github.com/technomancy/leiningen

import (
	"fmt"
	
	"os/exec"
)

func installLeiningen() {
	// Método 1: Descargar y extraer .tar.gz
	leiningen_tar_url := "https://github.com/technomancy/leiningen/archive/refs/tags/2.11.2.tar.gz"
	leiningen_cmd_tar := exec.Command("curl", "-L", leiningen_tar_url, "-o", "package.tar.gz")
	err := leiningen_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	leiningen_zip_url := "https://github.com/technomancy/leiningen/archive/refs/tags/2.11.2.zip"
	leiningen_cmd_zip := exec.Command("curl", "-L", leiningen_zip_url, "-o", "package.zip")
	err = leiningen_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	leiningen_bin_url := "https://github.com/technomancy/leiningen/archive/refs/tags/2.11.2.bin"
	leiningen_cmd_bin := exec.Command("curl", "-L", leiningen_bin_url, "-o", "binary.bin")
	err = leiningen_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	leiningen_src_url := "https://github.com/technomancy/leiningen/archive/refs/tags/2.11.2.src.tar.gz"
	leiningen_cmd_src := exec.Command("curl", "-L", leiningen_src_url, "-o", "source.tar.gz")
	err = leiningen_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	leiningen_cmd_direct := exec.Command("./binary")
	err = leiningen_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
exec.Command("latte", "install", "openjdk")
}

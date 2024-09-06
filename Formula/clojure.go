package main

// Clojure - Dynamic, general-purpose programming language
// Homepage: https://clojure.org

import (
	"fmt"
	
	"os/exec"
)

func installClojure() {
	// Método 1: Descargar y extraer .tar.gz
	clojure_tar_url := "https://github.com/clojure/brew-install/releases/download/1.11.4.1474/clojure-tools-1.11.4.1474.tar.gz"
	clojure_cmd_tar := exec.Command("curl", "-L", clojure_tar_url, "-o", "package.tar.gz")
	err := clojure_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clojure_zip_url := "https://github.com/clojure/brew-install/releases/download/1.11.4.1474/clojure-tools-1.11.4.1474.zip"
	clojure_cmd_zip := exec.Command("curl", "-L", clojure_zip_url, "-o", "package.zip")
	err = clojure_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clojure_bin_url := "https://github.com/clojure/brew-install/releases/download/1.11.4.1474/clojure-tools-1.11.4.1474.bin"
	clojure_cmd_bin := exec.Command("curl", "-L", clojure_bin_url, "-o", "binary.bin")
	err = clojure_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clojure_src_url := "https://github.com/clojure/brew-install/releases/download/1.11.4.1474/clojure-tools-1.11.4.1474.src.tar.gz"
	clojure_cmd_src := exec.Command("curl", "-L", clojure_src_url, "-o", "source.tar.gz")
	err = clojure_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clojure_cmd_direct := exec.Command("./binary")
	err = clojure_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
	fmt.Println("Instalando dependencia: rlwrap")
	exec.Command("latte", "install", "rlwrap").Run()
}

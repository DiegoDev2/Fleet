package main

// Clojurescript - Clojure to JS compiler
// Homepage: https://github.com/clojure/clojurescript

import (
	"fmt"
	
	"os/exec"
)

func installClojurescript() {
	// Método 1: Descargar y extraer .tar.gz
	clojurescript_tar_url := "https://github.com/clojure/clojurescript/releases/download/r1.11.132/cljs.jar"
	clojurescript_cmd_tar := exec.Command("curl", "-L", clojurescript_tar_url, "-o", "package.tar.gz")
	err := clojurescript_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	clojurescript_zip_url := "https://github.com/clojure/clojurescript/releases/download/r1.11.132/cljs.jar"
	clojurescript_cmd_zip := exec.Command("curl", "-L", clojurescript_zip_url, "-o", "package.zip")
	err = clojurescript_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	clojurescript_bin_url := "https://github.com/clojure/clojurescript/releases/download/r1.11.132/cljs.jar"
	clojurescript_cmd_bin := exec.Command("curl", "-L", clojurescript_bin_url, "-o", "binary.bin")
	err = clojurescript_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	clojurescript_src_url := "https://github.com/clojure/clojurescript/releases/download/r1.11.132/cljs.jar"
	clojurescript_cmd_src := exec.Command("curl", "-L", clojurescript_src_url, "-o", "source.tar.gz")
	err = clojurescript_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	clojurescript_cmd_direct := exec.Command("./binary")
	err = clojurescript_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openjdk")
	exec.Command("latte", "install", "openjdk").Run()
}

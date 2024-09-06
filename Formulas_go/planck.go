package main

// Planck - Stand-alone ClojureScript REPL
// Homepage: https://planck-repl.org/

import (
	"fmt"
	
	"os/exec"
)

func installPlanck() {
	// Método 1: Descargar y extraer .tar.gz
	planck_tar_url := "https://github.com/planck-repl/planck/archive/refs/tags/2.28.0.tar.gz"
	planck_cmd_tar := exec.Command("curl", "-L", planck_tar_url, "-o", "package.tar.gz")
	err := planck_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	planck_zip_url := "https://github.com/planck-repl/planck/archive/refs/tags/2.28.0.zip"
	planck_cmd_zip := exec.Command("curl", "-L", planck_zip_url, "-o", "package.zip")
	err = planck_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	planck_bin_url := "https://github.com/planck-repl/planck/archive/refs/tags/2.28.0.bin"
	planck_cmd_bin := exec.Command("curl", "-L", planck_bin_url, "-o", "binary.bin")
	err = planck_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	planck_src_url := "https://github.com/planck-repl/planck/archive/refs/tags/2.28.0.src.tar.gz"
	planck_cmd_src := exec.Command("curl", "-L", planck_src_url, "-o", "source.tar.gz")
	err = planck_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	planck_cmd_direct := exec.Command("./binary")
	err = planck_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: clojure")
exec.Command("latte", "install", "clojure")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libzip")
exec.Command("latte", "install", "libzip")
	fmt.Println("Instalando dependencia: webkitgtk")
exec.Command("latte", "install", "webkitgtk")
}

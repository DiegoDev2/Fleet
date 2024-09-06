package main

// QuickLintJs - Find bugs in your JavaScript code
// Homepage: https://quick-lint-js.com/

import (
	"fmt"
	
	"os/exec"
)

func installQuickLintJs() {
	// Método 1: Descargar y extraer .tar.gz
	quicklintjs_tar_url := "https://c.quick-lint-js.com/releases/3.2.0/source/quick-lint-js-3.2.0.tar.gz"
	quicklintjs_cmd_tar := exec.Command("curl", "-L", quicklintjs_tar_url, "-o", "package.tar.gz")
	err := quicklintjs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	quicklintjs_zip_url := "https://c.quick-lint-js.com/releases/3.2.0/source/quick-lint-js-3.2.0.zip"
	quicklintjs_cmd_zip := exec.Command("curl", "-L", quicklintjs_zip_url, "-o", "package.zip")
	err = quicklintjs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	quicklintjs_bin_url := "https://c.quick-lint-js.com/releases/3.2.0/source/quick-lint-js-3.2.0.bin"
	quicklintjs_cmd_bin := exec.Command("curl", "-L", quicklintjs_bin_url, "-o", "binary.bin")
	err = quicklintjs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	quicklintjs_src_url := "https://c.quick-lint-js.com/releases/3.2.0/source/quick-lint-js-3.2.0.src.tar.gz"
	quicklintjs_cmd_src := exec.Command("curl", "-L", quicklintjs_src_url, "-o", "source.tar.gz")
	err = quicklintjs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	quicklintjs_cmd_direct := exec.Command("./binary")
	err = quicklintjs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: googletest")
exec.Command("latte", "install", "googletest")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: simdjson")
exec.Command("latte", "install", "simdjson")
}

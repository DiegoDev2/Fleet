package main

// FluentBit - Fast and Lightweight Logs and Metrics processor
// Homepage: https://github.com/fluent/fluent-bit

import (
	"fmt"
	
	"os/exec"
)

func installFluentBit() {
	// Método 1: Descargar y extraer .tar.gz
	fluentbit_tar_url := "https://github.com/fluent/fluent-bit/archive/refs/tags/v3.1.7.tar.gz"
	fluentbit_cmd_tar := exec.Command("curl", "-L", fluentbit_tar_url, "-o", "package.tar.gz")
	err := fluentbit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fluentbit_zip_url := "https://github.com/fluent/fluent-bit/archive/refs/tags/v3.1.7.zip"
	fluentbit_cmd_zip := exec.Command("curl", "-L", fluentbit_zip_url, "-o", "package.zip")
	err = fluentbit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fluentbit_bin_url := "https://github.com/fluent/fluent-bit/archive/refs/tags/v3.1.7.bin"
	fluentbit_cmd_bin := exec.Command("curl", "-L", fluentbit_bin_url, "-o", "binary.bin")
	err = fluentbit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fluentbit_src_url := "https://github.com/fluent/fluent-bit/archive/refs/tags/v3.1.7.src.tar.gz"
	fluentbit_cmd_src := exec.Command("curl", "-L", fluentbit_src_url, "-o", "source.tar.gz")
	err = fluentbit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fluentbit_cmd_direct := exec.Command("./binary")
	err = fluentbit_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libyaml")
exec.Command("latte", "install", "libyaml")
	fmt.Println("Instalando dependencia: luajit")
exec.Command("latte", "install", "luajit")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}

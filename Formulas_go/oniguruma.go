package main

// Oniguruma - Regular expressions library
// Homepage: https://github.com/kkos/oniguruma/

import (
	"fmt"
	
	"os/exec"
)

func installOniguruma() {
	// Método 1: Descargar y extraer .tar.gz
	oniguruma_tar_url := "https://github.com/kkos/oniguruma/releases/download/v6.9.9/onig-6.9.9.tar.gz"
	oniguruma_cmd_tar := exec.Command("curl", "-L", oniguruma_tar_url, "-o", "package.tar.gz")
	err := oniguruma_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	oniguruma_zip_url := "https://github.com/kkos/oniguruma/releases/download/v6.9.9/onig-6.9.9.zip"
	oniguruma_cmd_zip := exec.Command("curl", "-L", oniguruma_zip_url, "-o", "package.zip")
	err = oniguruma_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	oniguruma_bin_url := "https://github.com/kkos/oniguruma/releases/download/v6.9.9/onig-6.9.9.bin"
	oniguruma_cmd_bin := exec.Command("curl", "-L", oniguruma_bin_url, "-o", "binary.bin")
	err = oniguruma_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	oniguruma_src_url := "https://github.com/kkos/oniguruma/releases/download/v6.9.9/onig-6.9.9.src.tar.gz"
	oniguruma_cmd_src := exec.Command("curl", "-L", oniguruma_src_url, "-o", "source.tar.gz")
	err = oniguruma_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	oniguruma_cmd_direct := exec.Command("./binary")
	err = oniguruma_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
}

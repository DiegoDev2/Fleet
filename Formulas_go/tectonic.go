package main

// Tectonic - Modernized, complete, self-contained TeX/LaTeX engine
// Homepage: https://tectonic-typesetting.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installTectonic() {
	// Método 1: Descargar y extraer .tar.gz
	tectonic_tar_url := "https://github.com/tectonic-typesetting/tectonic/archive/refs/tags/tectonic@0.15.0.tar.gz"
	tectonic_cmd_tar := exec.Command("curl", "-L", tectonic_tar_url, "-o", "package.tar.gz")
	err := tectonic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	tectonic_zip_url := "https://github.com/tectonic-typesetting/tectonic/archive/refs/tags/tectonic@0.15.0.zip"
	tectonic_cmd_zip := exec.Command("curl", "-L", tectonic_zip_url, "-o", "package.zip")
	err = tectonic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	tectonic_bin_url := "https://github.com/tectonic-typesetting/tectonic/archive/refs/tags/tectonic@0.15.0.bin"
	tectonic_cmd_bin := exec.Command("curl", "-L", tectonic_bin_url, "-o", "binary.bin")
	err = tectonic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	tectonic_src_url := "https://github.com/tectonic-typesetting/tectonic/archive/refs/tags/tectonic@0.15.0.src.tar.gz"
	tectonic_cmd_src := exec.Command("curl", "-L", tectonic_src_url, "-o", "source.tar.gz")
	err = tectonic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	tectonic_cmd_direct := exec.Command("./binary")
	err = tectonic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: rust")
exec.Command("latte", "install", "rust")
	fmt.Println("Instalando dependencia: freetype")
exec.Command("latte", "install", "freetype")
	fmt.Println("Instalando dependencia: graphite2")
exec.Command("latte", "install", "graphite2")
	fmt.Println("Instalando dependencia: harfbuzz")
exec.Command("latte", "install", "harfbuzz")
	fmt.Println("Instalando dependencia: icu4c")
exec.Command("latte", "install", "icu4c")
	fmt.Println("Instalando dependencia: libpng")
exec.Command("latte", "install", "libpng")
	fmt.Println("Instalando dependencia: openssl@3")
exec.Command("latte", "install", "openssl@3")
}

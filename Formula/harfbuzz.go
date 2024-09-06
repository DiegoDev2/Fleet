package main

// Harfbuzz - OpenType text shaping engine
// Homepage: https://github.com/harfbuzz/harfbuzz

import (
	"fmt"
	
	"os/exec"
)

func installHarfbuzz() {
	// Método 1: Descargar y extraer .tar.gz
	harfbuzz_tar_url := "https://github.com/harfbuzz/harfbuzz/archive/refs/tags/9.0.0.tar.gz"
	harfbuzz_cmd_tar := exec.Command("curl", "-L", harfbuzz_tar_url, "-o", "package.tar.gz")
	err := harfbuzz_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	harfbuzz_zip_url := "https://github.com/harfbuzz/harfbuzz/archive/refs/tags/9.0.0.zip"
	harfbuzz_cmd_zip := exec.Command("curl", "-L", harfbuzz_zip_url, "-o", "package.zip")
	err = harfbuzz_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	harfbuzz_bin_url := "https://github.com/harfbuzz/harfbuzz/archive/refs/tags/9.0.0.bin"
	harfbuzz_cmd_bin := exec.Command("curl", "-L", harfbuzz_bin_url, "-o", "binary.bin")
	err = harfbuzz_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	harfbuzz_src_url := "https://github.com/harfbuzz/harfbuzz/archive/refs/tags/9.0.0.src.tar.gz"
	harfbuzz_cmd_src := exec.Command("curl", "-L", harfbuzz_src_url, "-o", "source.tar.gz")
	err = harfbuzz_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	harfbuzz_cmd_direct := exec.Command("./binary")
	err = harfbuzz_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gobject-introspection")
	exec.Command("latte", "install", "gobject-introspection").Run()
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: pygobject3")
	exec.Command("latte", "install", "pygobject3").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: freetype")
	exec.Command("latte", "install", "freetype").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: graphite2")
	exec.Command("latte", "install", "graphite2").Run()
	fmt.Println("Instalando dependencia: icu4c")
	exec.Command("latte", "install", "icu4c").Run()
}

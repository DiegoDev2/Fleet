package main

// Lilv - C library to use LV2 plugins
// Homepage: https://drobilla.net/software/lilv.html

import (
	"fmt"
	
	"os/exec"
)

func installLilv() {
	// Método 1: Descargar y extraer .tar.gz
	lilv_tar_url := "https://download.drobilla.net/lilv-0.24.24.tar.xz"
	lilv_cmd_tar := exec.Command("curl", "-L", lilv_tar_url, "-o", "package.tar.gz")
	err := lilv_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lilv_zip_url := "https://download.drobilla.net/lilv-0.24.24.tar.xz"
	lilv_cmd_zip := exec.Command("curl", "-L", lilv_zip_url, "-o", "package.zip")
	err = lilv_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lilv_bin_url := "https://download.drobilla.net/lilv-0.24.24.tar.xz"
	lilv_cmd_bin := exec.Command("curl", "-L", lilv_bin_url, "-o", "binary.bin")
	err = lilv_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lilv_src_url := "https://download.drobilla.net/lilv-0.24.24.tar.xz"
	lilv_cmd_src := exec.Command("curl", "-L", lilv_src_url, "-o", "source.tar.gz")
	err = lilv_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lilv_cmd_direct := exec.Command("./binary")
	err = lilv_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
	exec.Command("latte", "install", "meson").Run()
	fmt.Println("Instalando dependencia: ninja")
	exec.Command("latte", "install", "ninja").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: libsndfile")
	exec.Command("latte", "install", "libsndfile").Run()
	fmt.Println("Instalando dependencia: lv2")
	exec.Command("latte", "install", "lv2").Run()
	fmt.Println("Instalando dependencia: serd")
	exec.Command("latte", "install", "serd").Run()
	fmt.Println("Instalando dependencia: sord")
	exec.Command("latte", "install", "sord").Run()
	fmt.Println("Instalando dependencia: sratom")
	exec.Command("latte", "install", "sratom").Run()
	fmt.Println("Instalando dependencia: zix")
	exec.Command("latte", "install", "zix").Run()
}

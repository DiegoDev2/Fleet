package main

// Uvw - Header-only, event based, tiny and easy to use libuv wrapper in modern C++
// Homepage: https://github.com/skypjack/uvw

import (
	"fmt"
	
	"os/exec"
)

func installUvw() {
	// Método 1: Descargar y extraer .tar.gz
	uvw_tar_url := "https://github.com/skypjack/uvw/archive/refs/tags/v3.4.0_libuv_v1.48.tar.gz"
	uvw_cmd_tar := exec.Command("curl", "-L", uvw_tar_url, "-o", "package.tar.gz")
	err := uvw_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	uvw_zip_url := "https://github.com/skypjack/uvw/archive/refs/tags/v3.4.0_libuv_v1.48.zip"
	uvw_cmd_zip := exec.Command("curl", "-L", uvw_zip_url, "-o", "package.zip")
	err = uvw_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	uvw_bin_url := "https://github.com/skypjack/uvw/archive/refs/tags/v3.4.0_libuv_v1.48.bin"
	uvw_cmd_bin := exec.Command("curl", "-L", uvw_bin_url, "-o", "binary.bin")
	err = uvw_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	uvw_src_url := "https://github.com/skypjack/uvw/archive/refs/tags/v3.4.0_libuv_v1.48.src.tar.gz"
	uvw_cmd_src := exec.Command("curl", "-L", uvw_src_url, "-o", "source.tar.gz")
	err = uvw_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	uvw_cmd_direct := exec.Command("./binary")
	err = uvw_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libuv")
exec.Command("latte", "install", "libuv")
}

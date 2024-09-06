package main

// Valabind - Vala bindings for radare, reverse engineering framework
// Homepage: https://github.com/radare/valabind

import (
	"fmt"
	
	"os/exec"
)

func installValabind() {
	// Método 1: Descargar y extraer .tar.gz
	valabind_tar_url := "https://github.com/radare/valabind/archive/refs/tags/1.8.0.tar.gz"
	valabind_cmd_tar := exec.Command("curl", "-L", valabind_tar_url, "-o", "package.tar.gz")
	err := valabind_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	valabind_zip_url := "https://github.com/radare/valabind/archive/refs/tags/1.8.0.zip"
	valabind_cmd_zip := exec.Command("curl", "-L", valabind_zip_url, "-o", "package.zip")
	err = valabind_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	valabind_bin_url := "https://github.com/radare/valabind/archive/refs/tags/1.8.0.bin"
	valabind_cmd_bin := exec.Command("curl", "-L", valabind_bin_url, "-o", "binary.bin")
	err = valabind_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	valabind_src_url := "https://github.com/radare/valabind/archive/refs/tags/1.8.0.src.tar.gz"
	valabind_cmd_src := exec.Command("curl", "-L", valabind_src_url, "-o", "source.tar.gz")
	err = valabind_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	valabind_cmd_direct := exec.Command("./binary")
	err = valabind_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: swig")
exec.Command("latte", "install", "swig")
	fmt.Println("Instalando dependencia: vala")
exec.Command("latte", "install", "vala")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
}

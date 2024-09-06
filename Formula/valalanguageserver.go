package main

// ValaLanguageServer - Code Intelligence for Vala & Genie
// Homepage: https://github.com/vala-lang/vala-language-server

import (
	"fmt"
	
	"os/exec"
)

func installValaLanguageServer() {
	// Método 1: Descargar y extraer .tar.gz
	valalanguageserver_tar_url := "https://github.com/vala-lang/vala-language-server/releases/download/0.48.7/vala-language-server-0.48.7.tar.xz"
	valalanguageserver_cmd_tar := exec.Command("curl", "-L", valalanguageserver_tar_url, "-o", "package.tar.gz")
	err := valalanguageserver_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	valalanguageserver_zip_url := "https://github.com/vala-lang/vala-language-server/releases/download/0.48.7/vala-language-server-0.48.7.tar.xz"
	valalanguageserver_cmd_zip := exec.Command("curl", "-L", valalanguageserver_zip_url, "-o", "package.zip")
	err = valalanguageserver_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	valalanguageserver_bin_url := "https://github.com/vala-lang/vala-language-server/releases/download/0.48.7/vala-language-server-0.48.7.tar.xz"
	valalanguageserver_cmd_bin := exec.Command("curl", "-L", valalanguageserver_bin_url, "-o", "binary.bin")
	err = valalanguageserver_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	valalanguageserver_src_url := "https://github.com/vala-lang/vala-language-server/releases/download/0.48.7/vala-language-server-0.48.7.tar.xz"
	valalanguageserver_cmd_src := exec.Command("curl", "-L", valalanguageserver_src_url, "-o", "source.tar.gz")
	err = valalanguageserver_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	valalanguageserver_cmd_direct := exec.Command("./binary")
	err = valalanguageserver_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: json-glib")
	exec.Command("latte", "install", "json-glib").Run()
	fmt.Println("Instalando dependencia: jsonrpc-glib")
	exec.Command("latte", "install", "jsonrpc-glib").Run()
	fmt.Println("Instalando dependencia: libgee")
	exec.Command("latte", "install", "libgee").Run()
	fmt.Println("Instalando dependencia: vala")
	exec.Command("latte", "install", "vala").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}

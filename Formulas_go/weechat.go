package main

// Weechat - Extensible IRC client
// Homepage: https://www.weechat.org

import (
	"fmt"
	
	"os/exec"
)

func installWeechat() {
	// Método 1: Descargar y extraer .tar.gz
	weechat_tar_url := "https://weechat.org/files/src/weechat-4.4.1.tar.xz"
	weechat_cmd_tar := exec.Command("curl", "-L", weechat_tar_url, "-o", "package.tar.gz")
	err := weechat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	weechat_zip_url := "https://weechat.org/files/src/weechat-4.4.1.tar.xz"
	weechat_cmd_zip := exec.Command("curl", "-L", weechat_zip_url, "-o", "package.zip")
	err = weechat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	weechat_bin_url := "https://weechat.org/files/src/weechat-4.4.1.tar.xz"
	weechat_cmd_bin := exec.Command("curl", "-L", weechat_bin_url, "-o", "binary.bin")
	err = weechat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	weechat_src_url := "https://weechat.org/files/src/weechat-4.4.1.tar.xz"
	weechat_cmd_src := exec.Command("curl", "-L", weechat_src_url, "-o", "source.tar.gz")
	err = weechat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	weechat_cmd_direct := exec.Command("./binary")
	err = weechat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: asciidoctor")
exec.Command("latte", "install", "asciidoctor")
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: aspell")
exec.Command("latte", "install", "aspell")
	fmt.Println("Instalando dependencia: cjson")
exec.Command("latte", "install", "cjson")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gnutls")
exec.Command("latte", "install", "gnutls")
	fmt.Println("Instalando dependencia: libgcrypt")
exec.Command("latte", "install", "libgcrypt")
	fmt.Println("Instalando dependencia: lua")
exec.Command("latte", "install", "lua")
	fmt.Println("Instalando dependencia: ncurses")
exec.Command("latte", "install", "ncurses")
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
	fmt.Println("Instalando dependencia: ruby")
exec.Command("latte", "install", "ruby")
	fmt.Println("Instalando dependencia: zstd")
exec.Command("latte", "install", "zstd")
	fmt.Println("Instalando dependencia: libgpg-error")
exec.Command("latte", "install", "libgpg-error")
}

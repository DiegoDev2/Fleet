package main

// JoplinCli - Note taking and to-do application with synchronization capabilities
// Homepage: https://joplinapp.org/

import (
	"fmt"
	
	"os/exec"
)

func installJoplinCli() {
	// Método 1: Descargar y extraer .tar.gz
	joplincli_tar_url := "https://registry.npmjs.org/joplin/-/joplin-3.0.1.tgz"
	joplincli_cmd_tar := exec.Command("curl", "-L", joplincli_tar_url, "-o", "package.tar.gz")
	err := joplincli_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	joplincli_zip_url := "https://registry.npmjs.org/joplin/-/joplin-3.0.1.tgz"
	joplincli_cmd_zip := exec.Command("curl", "-L", joplincli_zip_url, "-o", "package.zip")
	err = joplincli_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	joplincli_bin_url := "https://registry.npmjs.org/joplin/-/joplin-3.0.1.tgz"
	joplincli_cmd_bin := exec.Command("curl", "-L", joplincli_bin_url, "-o", "binary.bin")
	err = joplincli_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	joplincli_src_url := "https://registry.npmjs.org/joplin/-/joplin-3.0.1.tgz"
	joplincli_cmd_src := exec.Command("curl", "-L", joplincli_src_url, "-o", "source.tar.gz")
	err = joplincli_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	joplincli_cmd_direct := exec.Command("./binary")
	err = joplincli_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python-setuptools")
	exec.Command("latte", "install", "python-setuptools").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: node")
	exec.Command("latte", "install", "node").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: vips")
	exec.Command("latte", "install", "vips").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
	fmt.Println("Instalando dependencia: terminal-notifier")
	exec.Command("latte", "install", "terminal-notifier").Run()
	fmt.Println("Instalando dependencia: libsecret")
	exec.Command("latte", "install", "libsecret").Run()
}

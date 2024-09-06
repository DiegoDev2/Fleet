package main

// Graphene - Thin layer of graphic data types
// Homepage: https://ebassi.github.io/graphene/

import (
	"fmt"
	
	"os/exec"
)

func installGraphene() {
	// Método 1: Descargar y extraer .tar.gz
	graphene_tar_url := "https://github.com/ebassi/graphene/archive/refs/tags/1.10.8.tar.gz"
	graphene_cmd_tar := exec.Command("curl", "-L", graphene_tar_url, "-o", "package.tar.gz")
	err := graphene_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	graphene_zip_url := "https://github.com/ebassi/graphene/archive/refs/tags/1.10.8.zip"
	graphene_cmd_zip := exec.Command("curl", "-L", graphene_zip_url, "-o", "package.zip")
	err = graphene_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	graphene_bin_url := "https://github.com/ebassi/graphene/archive/refs/tags/1.10.8.bin"
	graphene_cmd_bin := exec.Command("curl", "-L", graphene_bin_url, "-o", "binary.bin")
	err = graphene_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	graphene_src_url := "https://github.com/ebassi/graphene/archive/refs/tags/1.10.8.src.tar.gz"
	graphene_cmd_src := exec.Command("curl", "-L", graphene_src_url, "-o", "source.tar.gz")
	err = graphene_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	graphene_cmd_direct := exec.Command("./binary")
	err = graphene_cmd_direct.Run()
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
	fmt.Println("Instalando dependencia: glib")
	exec.Command("latte", "install", "glib").Run()
	fmt.Println("Instalando dependencia: gettext")
	exec.Command("latte", "install", "gettext").Run()
}

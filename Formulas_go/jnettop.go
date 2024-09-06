package main

// Jnettop - View hosts/ports taking up the most network traffic
// Homepage: https://sourceforge.net/projects/jnettop/

import (
	"fmt"
	
	"os/exec"
)

func installJnettop() {
	// Método 1: Descargar y extraer .tar.gz
	jnettop_tar_url := "https://downloads.sourceforge.net/project/jnettop/jnettop/0.13/jnettop-0.13.0.tar.gz"
	jnettop_cmd_tar := exec.Command("curl", "-L", jnettop_tar_url, "-o", "package.tar.gz")
	err := jnettop_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jnettop_zip_url := "https://downloads.sourceforge.net/project/jnettop/jnettop/0.13/jnettop-0.13.0.zip"
	jnettop_cmd_zip := exec.Command("curl", "-L", jnettop_zip_url, "-o", "package.zip")
	err = jnettop_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jnettop_bin_url := "https://downloads.sourceforge.net/project/jnettop/jnettop/0.13/jnettop-0.13.0.bin"
	jnettop_cmd_bin := exec.Command("curl", "-L", jnettop_bin_url, "-o", "binary.bin")
	err = jnettop_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jnettop_src_url := "https://downloads.sourceforge.net/project/jnettop/jnettop/0.13/jnettop-0.13.0.src.tar.gz"
	jnettop_cmd_src := exec.Command("curl", "-L", jnettop_src_url, "-o", "source.tar.gz")
	err = jnettop_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jnettop_cmd_direct := exec.Command("./binary")
	err = jnettop_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: berkeley-db@5")
exec.Command("latte", "install", "berkeley-db@5")
}

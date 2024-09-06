package main

// Xauth - X.Org Applications: xauth
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installXauth() {
	// Método 1: Descargar y extraer .tar.gz
	xauth_tar_url := "https://www.x.org/pub/individual/app/xauth-1.1.3.tar.xz"
	xauth_cmd_tar := exec.Command("curl", "-L", xauth_tar_url, "-o", "package.tar.gz")
	err := xauth_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xauth_zip_url := "https://www.x.org/pub/individual/app/xauth-1.1.3.tar.xz"
	xauth_cmd_zip := exec.Command("curl", "-L", xauth_zip_url, "-o", "package.zip")
	err = xauth_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xauth_bin_url := "https://www.x.org/pub/individual/app/xauth-1.1.3.tar.xz"
	xauth_cmd_bin := exec.Command("curl", "-L", xauth_bin_url, "-o", "binary.bin")
	err = xauth_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xauth_src_url := "https://www.x.org/pub/individual/app/xauth-1.1.3.tar.xz"
	xauth_cmd_src := exec.Command("curl", "-L", xauth_src_url, "-o", "source.tar.gz")
	err = xauth_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xauth_cmd_direct := exec.Command("./binary")
	err = xauth_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: util-macros")
	exec.Command("latte", "install", "util-macros").Run()
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
	fmt.Println("Instalando dependencia: libxau")
	exec.Command("latte", "install", "libxau").Run()
	fmt.Println("Instalando dependencia: libxext")
	exec.Command("latte", "install", "libxext").Run()
	fmt.Println("Instalando dependencia: libxmu")
	exec.Command("latte", "install", "libxmu").Run()
	fmt.Println("Instalando dependencia: libxcb")
	exec.Command("latte", "install", "libxcb").Run()
	fmt.Println("Instalando dependencia: libxdmcp")
	exec.Command("latte", "install", "libxdmcp").Run()
}

package main

// Libxdmcp - X.Org: X Display Manager Control Protocol library
// Homepage: https://www.x.org/

import (
	"fmt"
	
	"os/exec"
)

func installLibxdmcp() {
	// Método 1: Descargar y extraer .tar.gz
	libxdmcp_tar_url := "https://www.x.org/archive/individual/lib/libXdmcp-1.1.5.tar.xz"
	libxdmcp_cmd_tar := exec.Command("curl", "-L", libxdmcp_tar_url, "-o", "package.tar.gz")
	err := libxdmcp_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libxdmcp_zip_url := "https://www.x.org/archive/individual/lib/libXdmcp-1.1.5.tar.xz"
	libxdmcp_cmd_zip := exec.Command("curl", "-L", libxdmcp_zip_url, "-o", "package.zip")
	err = libxdmcp_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libxdmcp_bin_url := "https://www.x.org/archive/individual/lib/libXdmcp-1.1.5.tar.xz"
	libxdmcp_cmd_bin := exec.Command("curl", "-L", libxdmcp_bin_url, "-o", "binary.bin")
	err = libxdmcp_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libxdmcp_src_url := "https://www.x.org/archive/individual/lib/libXdmcp-1.1.5.tar.xz"
	libxdmcp_cmd_src := exec.Command("curl", "-L", libxdmcp_src_url, "-o", "source.tar.gz")
	err = libxdmcp_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libxdmcp_cmd_direct := exec.Command("./binary")
	err = libxdmcp_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: xorgproto")
exec.Command("latte", "install", "xorgproto")
}

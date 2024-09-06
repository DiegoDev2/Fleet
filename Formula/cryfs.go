package main

// Cryfs - Encrypts your files so you can safely store them in Dropbox, iCloud, etc.
// Homepage: https://www.cryfs.org

import (
	"fmt"
	
	"os/exec"
)

func installCryfs() {
	// Método 1: Descargar y extraer .tar.gz
	cryfs_tar_url := "https://github.com/cryfs/cryfs/releases/download/0.11.4/cryfs-0.11.4.tar.gz"
	cryfs_cmd_tar := exec.Command("curl", "-L", cryfs_tar_url, "-o", "package.tar.gz")
	err := cryfs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cryfs_zip_url := "https://github.com/cryfs/cryfs/releases/download/0.11.4/cryfs-0.11.4.zip"
	cryfs_cmd_zip := exec.Command("curl", "-L", cryfs_zip_url, "-o", "package.zip")
	err = cryfs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cryfs_bin_url := "https://github.com/cryfs/cryfs/releases/download/0.11.4/cryfs-0.11.4.bin"
	cryfs_cmd_bin := exec.Command("curl", "-L", cryfs_bin_url, "-o", "binary.bin")
	err = cryfs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cryfs_src_url := "https://github.com/cryfs/cryfs/releases/download/0.11.4/cryfs-0.11.4.src.tar.gz"
	cryfs_cmd_src := exec.Command("curl", "-L", cryfs_src_url, "-o", "source.tar.gz")
	err = cryfs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cryfs_cmd_direct := exec.Command("./binary")
	err = cryfs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
	exec.Command("latte", "install", "cmake").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
	fmt.Println("Instalando dependencia: boost")
	exec.Command("latte", "install", "boost").Run()
	fmt.Println("Instalando dependencia: curl")
	exec.Command("latte", "install", "curl").Run()
	fmt.Println("Instalando dependencia: fmt")
	exec.Command("latte", "install", "fmt").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
	fmt.Println("Instalando dependencia: range-v3")
	exec.Command("latte", "install", "range-v3").Run()
	fmt.Println("Instalando dependencia: spdlog")
	exec.Command("latte", "install", "spdlog").Run()
}

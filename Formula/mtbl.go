package main

// Mtbl - Immutable sorted string table library
// Homepage: https://github.com/farsightsec/mtbl

import (
	"fmt"
	
	"os/exec"
)

func installMtbl() {
	// Método 1: Descargar y extraer .tar.gz
	mtbl_tar_url := "https://dl.farsightsecurity.com/dist/mtbl/mtbl-1.6.1.tar.gz"
	mtbl_cmd_tar := exec.Command("curl", "-L", mtbl_tar_url, "-o", "package.tar.gz")
	err := mtbl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mtbl_zip_url := "https://dl.farsightsecurity.com/dist/mtbl/mtbl-1.6.1.zip"
	mtbl_cmd_zip := exec.Command("curl", "-L", mtbl_zip_url, "-o", "package.zip")
	err = mtbl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mtbl_bin_url := "https://dl.farsightsecurity.com/dist/mtbl/mtbl-1.6.1.bin"
	mtbl_cmd_bin := exec.Command("curl", "-L", mtbl_bin_url, "-o", "binary.bin")
	err = mtbl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mtbl_src_url := "https://dl.farsightsecurity.com/dist/mtbl/mtbl-1.6.1.src.tar.gz"
	mtbl_cmd_src := exec.Command("curl", "-L", mtbl_src_url, "-o", "source.tar.gz")
	err = mtbl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mtbl_cmd_direct := exec.Command("./binary")
	err = mtbl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: snappy")
	exec.Command("latte", "install", "snappy").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}

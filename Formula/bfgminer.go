package main

// Bfgminer - Modular CPU/GPU/ASIC/FPGA miner written in C
// Homepage: https://web.archive.org/web/20221230131107/http://bfgminer.org/

import (
	"fmt"
	
	"os/exec"
)

func installBfgminer() {
	// Método 1: Descargar y extraer .tar.gz
	bfgminer_tar_url := "https://web.archive.org/web/20190824104403/http://bfgminer.org/files/latest/bfgminer-5.5.0.txz"
	bfgminer_cmd_tar := exec.Command("curl", "-L", bfgminer_tar_url, "-o", "package.tar.gz")
	err := bfgminer_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bfgminer_zip_url := "https://web.archive.org/web/20190824104403/http://bfgminer.org/files/latest/bfgminer-5.5.0.txz"
	bfgminer_cmd_zip := exec.Command("curl", "-L", bfgminer_zip_url, "-o", "package.zip")
	err = bfgminer_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bfgminer_bin_url := "https://web.archive.org/web/20190824104403/http://bfgminer.org/files/latest/bfgminer-5.5.0.txz"
	bfgminer_cmd_bin := exec.Command("curl", "-L", bfgminer_bin_url, "-o", "binary.bin")
	err = bfgminer_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bfgminer_src_url := "https://web.archive.org/web/20190824104403/http://bfgminer.org/files/latest/bfgminer-5.5.0.txz"
	bfgminer_cmd_src := exec.Command("curl", "-L", bfgminer_src_url, "-o", "source.tar.gz")
	err = bfgminer_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bfgminer_cmd_direct := exec.Command("./binary")
	err = bfgminer_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: hidapi")
	exec.Command("latte", "install", "hidapi").Run()
	fmt.Println("Instalando dependencia: libgcrypt")
	exec.Command("latte", "install", "libgcrypt").Run()
	fmt.Println("Instalando dependencia: libscrypt")
	exec.Command("latte", "install", "libscrypt").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: uthash")
	exec.Command("latte", "install", "uthash").Run()
	fmt.Println("Instalando dependencia: jansson")
	exec.Command("latte", "install", "jansson").Run()
	fmt.Println("Instalando dependencia: libevent")
	exec.Command("latte", "install", "libevent").Run()
	fmt.Println("Instalando dependencia: libusb")
	exec.Command("latte", "install", "libusb").Run()
}

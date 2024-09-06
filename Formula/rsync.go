package main

// Rsync - Utility that provides fast incremental file transfer
// Homepage: https://rsync.samba.org/

import (
	"fmt"
	
	"os/exec"
)

func installRsync() {
	// Método 1: Descargar y extraer .tar.gz
	rsync_tar_url := "https://rsync.samba.org/ftp/rsync/rsync-3.3.0.tar.gz"
	rsync_cmd_tar := exec.Command("curl", "-L", rsync_tar_url, "-o", "package.tar.gz")
	err := rsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	rsync_zip_url := "https://rsync.samba.org/ftp/rsync/rsync-3.3.0.zip"
	rsync_cmd_zip := exec.Command("curl", "-L", rsync_zip_url, "-o", "package.zip")
	err = rsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	rsync_bin_url := "https://rsync.samba.org/ftp/rsync/rsync-3.3.0.bin"
	rsync_cmd_bin := exec.Command("curl", "-L", rsync_bin_url, "-o", "binary.bin")
	err = rsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	rsync_src_url := "https://rsync.samba.org/ftp/rsync/rsync-3.3.0.src.tar.gz"
	rsync_cmd_src := exec.Command("curl", "-L", rsync_src_url, "-o", "source.tar.gz")
	err = rsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	rsync_cmd_direct := exec.Command("./binary")
	err = rsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: lz4")
	exec.Command("latte", "install", "lz4").Run()
	fmt.Println("Instalando dependencia: openssl@3")
	exec.Command("latte", "install", "openssl@3").Run()
	fmt.Println("Instalando dependencia: popt")
	exec.Command("latte", "install", "popt").Run()
	fmt.Println("Instalando dependencia: xxhash")
	exec.Command("latte", "install", "xxhash").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
